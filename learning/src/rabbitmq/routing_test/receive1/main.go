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
		log.Fatal("route mq conn err", err)
	}
	defer mq.Destory()

	queue, err := mq.DeclareQueue(&rabbitmq.QueueConfig{QueueName: "routingmq1"})
	if err != nil {
		log.Fatal("route mq declare queue err", err)
	}

	err = mq.DeclareRoutingExchange(&rabbitmq.ExchangeConfig{
		ExchangeName: "logs_direct",
	})
	if err != nil {
		log.Fatal("route mq declare fanout exchange err", err)
	}

	mq.BindRoutingQueue(&rabbitmq.QueueBindConfig{
		QueueName:    queue.Name,
		ExchangeName: "logs_direct",
		RoutingKey:   "route1",
	})

	mq.SubscribeRoutingMessage(&queue, &rabbitmq.ReceiveConfig{AutoAck: true}, func(message amqp.Delivery) {
		fmt.Println("receiver1 message", string(message.Body))
	})
}
