package main

// 消费者代码 work.go

import (
	"log"
	"os"

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

	// 创建通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明交换机
	err = ch.ExchangeDeclare(
		"logs_direct", // name 交换机名称
		"direct",      // type 交换机类型
		true,          // durable 持久化
		false,         // auto-deleted 当所有与交换机绑定的队列都被删除时，交换机会自动删除。如果设置为 false，交换机不会自动删除。
		false,         // internal 如果设置为 true，交换机将是内部的，这意味着它不能被客户端直接使用，只能通过其他交换机进行路由。通常设置为 false
		false,         // no-wait ：如果设置为 true，表示不等待服务器返回结果。交换机声明后，立即继续执行后续代码，不等待确认。如果设置为 false，程序会等待服务器的确认，确保交换机成功声明
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// 声明队列
	q, err := ch.QueueDeclare(
		"",    // name 队列名称
		false, // durable 持久化
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_direct", s)
		// 绑定交换机
		err = ch.QueueBind(
			q.Name,        // queue name 队列名
			s,             // routing key
			"logs_direct", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	// 接受消息
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
