package book

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"mq/lib"
	"mq/lib/errors"

	"github.com/streadway/amqp"
)

type ConsumerBody struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Consumer struct {
	Connection lib.IConnection
}

func NewConsumer(connection lib.IConnection) *Consumer {
	return &Consumer{
		Connection: connection,
	}
}

func (s *Consumer) Consume() {
	ch, err := s.Connection.GetMQConnection().Channel()
	if err != nil {
		log.Print(err)
		return
	}

	messages, err := ch.Consume(
		lib.RoutingKeyBook,
		lib.GetGenericName(lib.RoutingKeyBook),
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Print(err)
		return
	}

	for message := range messages {
		s.Process(&message)
	}
}

func (s *Consumer) Process(d *amqp.Delivery) {
	var body ConsumerBody

	err := json.Unmarshal(d.Body, &body)
	if err != nil {
		log.Print(err)

		err = d.Ack(false)
		if err != nil {
			log.Print(err)
		}

		return
	}

	err = s.Run(body)
	if err != nil {
		log.Println(err)
		err = d.Nack(false, true)
		if err != nil {
			log.Print(err)
		}
		return
	}

	err = d.Ack(false)
	if err != nil {
		log.Print(err)
	}
}

func (s *Consumer) Run(body any) error {
	form, ok := body.(ConsumerBody)
	if !ok {
		return errors.ErrTypeMismatch
	}

	durationTime := time.Second * 5
	time.Sleep(durationTime)

	fmt.Printf("ID: %s;\nName: %s", form.ID, form.Name)

	return nil
}
