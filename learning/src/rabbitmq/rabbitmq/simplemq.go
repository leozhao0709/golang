package rabbitmq

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
)

// SimpleMQ a simple mode rabbit mq
type SimpleMQ struct {
	rabbitMQ
}

// NewSimpleMq create a new simple mq
func NewSimpleMq(host string, port string, username string, password string, queueName string) (*SimpleMQ, error) {
	var err error
	var simpleMq = &SimpleMQ{}
	simpleMq.conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port))
	if err != nil {
		return nil, err
	}

	simpleMq.channel, err = simpleMq.conn.Channel()
	if err != nil {
		return nil, err
	}

	simpleMq.queue, err = simpleMq.channel.QueueDeclare(
		queueName, // queuename
		false,     // durable persist data or not
		false,     // autoDelete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)

	return simpleMq, nil
}

// Publish simple mq publish message
func (mq SimpleMQ) Publish(message amqp.Publishing) error {
	return mq.channel.Publish(
		"",            // exchange
		mq.queue.Name, // routing key
		false,         // false
		false,         // immediate
		message,       // publishing message
	)
}

// Consume simple mq consume message
func (mq SimpleMQ) Consume(receiveHandler func(message amqp.Delivery)) error {
	msgs, err := mq.channel.Consume(
		mq.queue.Name,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			receiveHandler(d)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}
