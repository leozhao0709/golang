package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Config RabbitMQ config
type Config struct {
	Host, Username, Password string
	Port                     int
}

// QueueConfig RabbiMQ queue configuration
// used for declare a queue
type QueueConfig struct {
	QueueName                              string
	Durable, AutoDelete, Exclusive, NoWait bool
	Args                                   amqp.Table
}

// PublishConfig publish message config
type PublishConfig struct {
	ExchangeName, Key    string
	Mandatory, Immediate bool
}

// ReceiveConfig receive message config
type ReceiveConfig struct {
	Consumer                            string
	AutoAck, Exclusive, NoLocal, NoWait bool
	Args                                amqp.Table
}

// ExchangeConfig RabbitMQ Exchange config
type ExchangeConfig struct {
	Durable, AutoDelete, Internal, NoWait bool
	Args                                  amqp.Table
}

// RabbitMQ raibbitmq struct
type RabbitMQ struct {
	URL     string
	Channel *amqp.Channel
	Conn    *amqp.Connection
}

// Destory close all the connection or channel
func (r *RabbitMQ) Destory() {
	r.Channel.Close()
	r.Conn.Close()
}

// New create a new Rabbit MQ
func New(config *Config) (*RabbitMQ, error) {
	var err error
	mq := &RabbitMQ{}

	if config == nil {
		config = &Config{}
	}

	if config.Host == "" {
		config.Host = "localhost"
	}

	if config.Port == 0 {
		config.Port = 5672
	}

	if config.Username == "" {
		config.Username = "guest"
	}

	if config.Password == "" {
		config.Password = "guest"
	}

	mq.Conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port))
	if err != nil {
		return nil, err
	}

	mq.Channel, err = mq.Conn.Channel()
	if err != nil {
		return nil, err
	}

	return mq, nil
}

// DeclareQueue declare the queue which are going to use
func (r *RabbitMQ) DeclareQueue(queueName string, extraConfig *QueueConfig) (amqp.Queue, error) {

	if extraConfig != nil {
		return r.Channel.QueueDeclare(
			queueName,              // queuename
			extraConfig.Durable,    // durable persist data or not
			extraConfig.AutoDelete, // autoDelete
			extraConfig.Exclusive,  // exclusive
			extraConfig.NoWait,     // no wait
			extraConfig.Args,       // arguments
		)
	}

	return r.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
}

// DeclareExchange declare the exchange
func (r *RabbitMQ) declareExchange(exchangeName, exchangeType string, extraExchangeConfig *ExchangeConfig) error {
	if extraExchangeConfig != nil {
		return r.Channel.ExchangeDeclare(
			exchangeName,
			exchangeType,
			extraExchangeConfig.Durable,
			extraExchangeConfig.AutoDelete,
			extraExchangeConfig.Internal,
			extraExchangeConfig.NoWait,
			extraExchangeConfig.Args,
		)
	}

	return r.Channel.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

// publishMessage publish a message
func (r *RabbitMQ) publishMessage(config *PublishConfig, message amqp.Publishing) error {
	return r.Channel.Publish(
		config.ExchangeName,
		config.Key,
		config.Mandatory,
		config.Immediate,
		message,
	)
}
