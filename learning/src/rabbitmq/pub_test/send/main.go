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
		log.Fatal("pub mq conn err", err)
	}
	defer mq.Destory()

	mq.DeclarePublishExchange(&rabbitmq.ExchangeConfig{
		ExchangeName: "logs",
	})

	for i := 0; i < 100; i++ {
		mq.SendPublishMessage(&rabbitmq.PublishConfig{
			ExchangeName: "logs",
			RoutingKey:   "",
		}, amqp.Publishing{ContentType: "text/plain", Body: []byte(fmt.Sprint("message", i))})
		time.Sleep(time.Second)
	}
}
