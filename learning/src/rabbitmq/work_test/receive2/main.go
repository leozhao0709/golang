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
		log.Fatal("work mq conn err", err)
	}
	defer mq.Destory()

	queue, err := mq.DeclareQueue(&rabbitmq.QueueConfig{QueueName: "workmq"})
	if err != nil {
		log.Fatal("work mq declare queue err", err)
	}

	mq.ReceiveWorkMessage(&queue, nil, func(message amqp.Delivery) {
		fmt.Println("receiver2 message", string(message.Body))
	})
}
