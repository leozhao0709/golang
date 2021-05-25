package main

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/musings/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	mq, err := rabbitmq.New(&rabbitmq.Config{})
	if err != nil {
		log.Fatal("Topic mq conn err", err)
	}
	defer mq.Destory()

	mq.DeclareTopicExchange(&rabbitmq.ExchangeConfig{
		ExchangeName: "logs_topic",
	})

	for i := 0; i < 100; i++ {
		mq.SendTopicMessage(&rabbitmq.PublishConfig{
			ExchangeName: "logs_topic",
			RoutingKey:   "lei.topic.one",
		}, amqp.Publishing{ContentType: "text/plain", Body: []byte(fmt.Sprint("message topic1 ", i))})
		mq.SendTopicMessage(&rabbitmq.PublishConfig{
			ExchangeName: "logs_topic",
			RoutingKey:   "lei.topic_test.two",
		}, amqp.Publishing{ContentType: "text/plain", Body: []byte(fmt.Sprint("message topic2 ", i))})
		time.Sleep(time.Second)
	}
}
