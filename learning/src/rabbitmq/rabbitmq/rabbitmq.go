package rabbitmq

import (
	"github.com/streadway/amqp"
)

// rabbitMQ raibbitmq struct
type rabbitMQ struct {
	url     string
	channel *amqp.Channel
	conn    *amqp.Connection
	queue   amqp.Queue
}

// Destory close all the connection or channel
func (r rabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}
