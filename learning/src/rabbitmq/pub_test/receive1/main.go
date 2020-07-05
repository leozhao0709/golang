package main

import (
	"fmt"
	"log"

	"github.com/leozhao0709/go-musings/src/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	mq, err := rabbitmq.New(&rabbitmq.Config{})
	if err != nil {
		log.Fatal("pub mq conn err", err)
	}
	defer mq.Destory()

	queue, err := mq.DeclareQueue(&rabbitmq.QueueConfig{QueueName: "pubmq1"})
	if err != nil {
		log.Fatal("pub mq declare queue err", err)
	}

	err = mq.DeclarePublishExchange(&rabbitmq.ExchangeConfig{
		ExchangeName: "logs",
	})
	if err != nil {
		log.Fatal("pub mq declare fanout exchange err", err)
	}

	mq.BindPublishQueue(&rabbitmq.QueueBindConfig{
		QueueName:    queue.Name,
		ExchangeName: "logs",
	})

	mq.SubscribePublishMessage(&queue, &rabbitmq.ReceiveConfig{AutoAck: true}, func(message amqp.Delivery) {
		fmt.Println("receiver2 message", string(message.Body))
	})
}
