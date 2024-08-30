// 生产者代码 send.go
package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 闯进通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明交换机
	err = ch.ExchangeDeclare(
		"logs_direct", // name 交换机名称
		"direct",      // type 交换机类型 直连交换机
		true,          // durable
		false,         // auto-deleted 当所有与交换机绑定的队列都被删除时，交换机会自动删除。如果设置为 false，交换机不会自动删除。
		false,         // internal 如果设置为 true，交换机将是内部的，这意味着它不能被客户端直接使用，只能通过其他交换机进行路由。通常设置为 false
		false,         // no-wait ：如果设置为 true，表示不等待服务器返回结果。交换机声明后，立即继续执行后续代码，不等待确认。如果设置为 false，程序会等待服务器的确认，确保交换机成功声明
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := bodyFrom(os.Args)
	err = ch.PublishWithContext(ctx,
		"logs_direct",         // exchange
		severityFrom(os.Args), // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
