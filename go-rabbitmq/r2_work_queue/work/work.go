package main

//消费者代码 work.go

import (
	"bytes"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count 表示在消费者发送 ack 确认之前，RabbitMQ 允许同一消费者最多接收多少条未确认的消息
		0,     // prefetch size 以字节为单位，指定消费者在发送 ack 之前可以接收的消息总大小。这里设置为 0，表示没有限制
		false, // global 如果设置为 true，则 prefetch count 和 prefetch size 将对整个通道（所有消费者）生效。如果设置为 false，则仅对当前消费者生效
	)
	failOnError(err, "Failed to set QoS")

	// 注册消费者
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte(".")) // 统计点数
			t := time.Duration(dotCount)                 // 计算睡眠时间
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false) // 确认消息，向RabbitMQ发送一个确认信号，表示该消息已被成功处理
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
