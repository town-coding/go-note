package main

// 生产者 send.go
import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strings"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// 建立连接
	conn, err := amqp.Dial("amqp://guest:guest@139.159.151.30:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 打开通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明交换机
	err = ch.ExchangeDeclare(
		"logs_topic", // name 交换机名称
		"topic",      // type 交换机类型 topic 主题模式
		true,         // durable 持久化
		false,        // auto-deleted 当所有与交换机绑定的队列都被删除时，交换机会自动删除。如果设置为 false，交换机不会自动删除。
		false,        // internal 如果设置为 true，交换机将是内部的，这意味着它不能被客户端直接使用，只能通过其他交换机进行路由。通常设置为 false
		false,        // no-wait ：如果设置为 true，表示不等待服务器返回结果。交换机声明后，立即继续执行后续代码，不等待确认。如果设置为 false，程序会等待服务器的确认，确保交换机成功声明
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := bodyFrom(os.Args)
	err = ch.PublishWithContext(ctx,
		"logs_topic",          // exchange
		severityFrom(os.Args), // routing key
		false,                 // mandatory 如果设置为 true，则表示如果消息无法路由到任何队列，RabbitMQ 将返回该消息给发布者。如果设置为 false，消息将被丢弃
		false,                 // immediate 如果设置为 true，表示如果消息发送时，队列中没有消费者，消息将不会入队列，并返回给发布者。一般情况下，这个选项很少使用，通常设置为 false
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
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}
