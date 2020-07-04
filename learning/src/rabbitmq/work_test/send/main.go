package main

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/learning/src/rabbitmq/rabbitmq"
	"github.com/streadway/amqp"
)

func main() {
	mq, err := rabbitmq.New(&rabbitmq.Config{})
	if err != nil {
		log.Fatal("work mq conn err", err)
	}
	defer mq.Destory()

	queue, err := mq.DeclareQueue("workmq", nil)
	if err != nil {
		log.Fatal("work mq declare queue err", err)
	}

	for i := 0; i < 100; i++ {
		mq.SendWorkMessage(queue, nil, amqp.Publishing{ContentType: "text/plain", Body: []byte(fmt.Sprint("message", i))})
		time.Sleep(time.Second)
	}
}
