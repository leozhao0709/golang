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
		log.Fatal("routing mq conn err", err)
	}
	defer mq.Destory()

	mq.DeclareRoutingExchange(&rabbitmq.ExchangeConfig{
		ExchangeName: "logs_direct",
	})

	for i := 0; i < 100; i++ {
		mq.SendRoutingMessage(&rabbitmq.PublishConfig{
			ExchangeName: "logs_direct",
			RoutingKey:   "route1",
		}, amqp.Publishing{ContentType: "text/plain", Body: []byte(fmt.Sprint("message", i))})
		mq.SendRoutingMessage(&rabbitmq.PublishConfig{
			ExchangeName: "logs_direct",
			RoutingKey:   "route2",
		}, amqp.Publishing{ContentType: "text/plain", Body: []byte(fmt.Sprint("message", i))})
		time.Sleep(time.Second)
	}
}
