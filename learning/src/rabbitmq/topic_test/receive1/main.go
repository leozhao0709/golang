package main

import (
	"fmt"
	"log"

	"github.com/leozhao0709/musings/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	mq, err := rabbitmq.New(&rabbitmq.Config{})
	if err != nil {
		log.Fatal("topic mq conn err", err)
	}
	defer mq.Destory()

	queue, err := mq.DeclareQueue(&rabbitmq.QueueConfig{QueueName: ""})
	if err != nil {
		log.Fatal("topic mq declare queue err", err)
	}

	err = mq.DeclareTopicExchange(&rabbitmq.ExchangeConfig{
		ExchangeName: "logs_topic",
	})
	if err != nil {
		log.Fatal("topic mq declare topic exchange err ", err)
	}

	err = mq.BindTopicQueue(&rabbitmq.QueueBindConfig{
		QueueName:    queue.Name,
		ExchangeName: "logs_topic",
		RoutingKey:   "#", // match all
	})

	if err != nil {
		log.Fatal("topic mq binding err", err)
	}

	err = mq.SubscribeTopicMessage(&queue, &rabbitmq.ReceiveConfig{AutoAck: true}, func(message amqp.Delivery) {
		fmt.Println("receiver1 message", string(message.Body))
	})

	if err != nil {
		log.Fatal("topic mq subscribe err", err)
	}
}
