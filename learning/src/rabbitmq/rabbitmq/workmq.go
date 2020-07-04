package rabbitmq

import (
	"github.com/streadway/amqp"
)

// SendWorkMessage send simple work message
func (r *RabbitMQ) SendWorkMessage(queue amqp.Queue, extraPublishConfig *PublishConfig, message amqp.Publishing) error {

	if extraPublishConfig != nil {
		extraPublishConfig.Key = queue.Name
		return r.publishMessage(
			extraPublishConfig,
			message,
		)
	}

	return r.Channel.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // Mandatory
		false,      // immediate
		message,    // publishing message
	)
}

// ReceiveWorkMessage receive simple work message
func (r *RabbitMQ) ReceiveWorkMessage(queue amqp.Queue, extraReceiveConfig *ReceiveConfig, receiveHandler func(message amqp.Delivery)) error {
	var msgs <-chan amqp.Delivery
	var err error

	if extraReceiveConfig != nil {
		msgs, err = r.Channel.Consume(
			queue.Name,
			extraReceiveConfig.Consumer,  // consumer
			extraReceiveConfig.AutoAck,   // auto-ack
			extraReceiveConfig.Exclusive, // exclusive
			extraReceiveConfig.NoLocal,   // no-local
			extraReceiveConfig.NoWait,    // no-wait
			extraReceiveConfig.Args,      // args
		)
	} else {
		msgs, err = r.Channel.Consume(
			queue.Name,
			"",    // consumer
			true,  // auto-ack
			false, // exclusive
			false, // no-local
			false, // no-wait
			nil,   // args
		)
	}

	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			receiveHandler(d)
		}
	}()

	<-forever
	return nil
}
