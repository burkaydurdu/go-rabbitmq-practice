package lib

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type IConnection interface {
	GetMQConnection() *amqp.Connection
}

type Connection struct {
	MQConnection *amqp.Connection
}

func NewConnection() *Connection {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		"guest", "guest", "localhost", "5672",
	)

	config := amqp.Config{
		Heartbeat: 30 * time.Second,
		Locale:    "en_US",
	}

	connection, err := amqp.DialConfig(url, config)
	if err != nil {
		panic(err)
	}

	return &Connection{
		MQConnection: connection,
	}
}

func (c *Connection) GetMQConnection() *amqp.Connection {
	if c.MQConnection.IsClosed() {
		panic("RabbitMQ Connection Is Closed")
	}
	return c.MQConnection
}
