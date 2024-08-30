package main

// 发布者代码 publish.go

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 打开通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明交换机
	err = ch.ExchangeDeclare(
		"logs",   // name 交换机名称
		"fanout", // type 交换机类型
		true,     // durable 持久化表示
		false,    // auto-deleted 当所有与交换机绑定的队列都被删除时，交换机会自动删除。如果设置为 false，交换机不会自动删除。
		false,    // internal 如果设置为 true，交换机将是内部的，这意味着它不能被客户端直接使用，只能通过其他交换机进行路由。通常设置为 false
		false,    // no-wait ：如果设置为 true，表示不等待服务器返回结果。交换机声明后，立即继续执行后续代码，不等待确认。如果设置为 false，程序会等待服务器的确认，确保交换机成功声明
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := bodyFrom(os.Args)
	// 用于将消息发布到指定的交换机
	err = ch.PublishWithContext(ctx,
		"logs", // exchange 交换机名称
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

// 获取启动命令行参数
func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
