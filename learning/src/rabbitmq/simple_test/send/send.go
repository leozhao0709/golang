package main

import (
	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/learning/src/rabbitmq/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	simpleMq, err := rabbitmq.NewSimpleMq("localhost", "5672", "guest", "guest", "simpleQueue")

	if err != nil {
		log.Fatal("simple mq conn err", err)
	}
	defer simpleMq.Destory()

	simpleMq.Publish(amqp.Publishing{ContentType: "text/plain", Body: []byte("Hello World!")})
}
