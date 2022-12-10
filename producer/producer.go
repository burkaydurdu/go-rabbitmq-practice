package producer

import (
	"fmt"
	"log"

	"mq/lib"

	"github.com/streadway/amqp"
)

func Publish(connection lib.IConnection, exchange, exchangeType, routingKey string, body []byte, priority uint8) error {
	if connection.GetMQConnection().IsClosed() {
		panic("RabbitMQ Connection is Closed")
	}
	connection.GetMQConnection()

	channel, err := connection.GetMQConnection().Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}

	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {
			log.Println(err)
		}
	}(channel)

	if err := channel.ExchangeDeclare(
		exchange,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	if err := channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Persistent,
			Priority:        priority,
		},
	); err != nil {
		return err
	}

	return nil
}
