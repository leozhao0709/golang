package rabbitmq

import "github.com/streadway/amqp"

// SendPublishMessage send a publish message
func (r *RabbitMQ) SendPublishMessage(queue amqp.Queue, exchangeName string, extraExchangeConfig *ExchangeConfig, extraPublishConfig *PublishConfig, message amqp.Publishing) error {
	err := r.declareExchange(
		exchangeName,
		"fanout",
		extraExchangeConfig,
	)

	if err != nil {
		return err
	}

	return r.publishMessage(
		extraPublishConfig,
		message,
	)
}
