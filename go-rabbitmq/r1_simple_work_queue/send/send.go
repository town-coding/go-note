package main

// 生产者代码 send.go

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

// failOnError 处理错误使用，后续不再描述
func failOnError(err error, msg string) {
	if err != nil {
		// 打印 错误日志
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// 1、建立mq连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2、建立通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个队列
	q, err := ch.QueueDeclare(
		"hello", // name 队列名称
		false,   // durable 是否持久化
		false,   // delete when unused  队列不再使用就删除
		false,   // exclusive 队列是否仅供当前连接使用
		false,   // no-wait 声明队列时是否等待RabbitMQ确认
		nil,     // arguments 传递额外的队列参数
	)
	failOnError(err, "Failed to declare a queue")
	// 创建一个带有超时的上下文，设置超时时间为5秒。在发送消息时，如果操作超时，上下文会取消操作
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 消息体
	body := "Hello World!"
	// 用于将消息发布到指定的交换机
	err = ch.PublishWithContext(ctx,
		"",     // exchange 交换机
		q.Name, // routing key 路由键
		false,  // mandatory 消息是否必须路由到队列
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
