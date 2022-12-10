package book

import (
	"encoding/json"
	"mq/lib"
	"mq/producer"
)

type ProducerBody struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type IProducer interface {
	Send(ID, Name string) error
}

type Producer struct {
	Connection lib.IConnection
}

func NewProducer(connection lib.IConnection) *Producer {
	return &Producer{
		Connection: connection,
	}
}

func (p *Producer) Send(id, name string) error {
	body := ProducerBody{
		ID:   id,
		Name: name,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return producer.Publish(p.Connection, lib.ProjectExchange, lib.ExchangeTypeTopic, lib.RoutingKeyBook, jsonBody, 0)
}
