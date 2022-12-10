package consumer

import (
	"log"

	"mq/lib"
	"mq/lib/errors"

	"github.com/streadway/amqp"
)

type IConsumer interface {
	Consume()
	Process(d *amqp.Delivery)
	Run(body any) error
}

type IService interface {
	Consume(workerType lib.WorkerType) error
}

type Service struct {
	BookConsumer IConsumer
}

func NewConsumeService(bookConsumer IConsumer) *Service {
	return &Service{
		BookConsumer: bookConsumer,
	}
}

func (s *Service) Consume(workerType lib.WorkerType) error {
	consumers, err := s.getConsumers(workerType)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	for _, consumer := range consumers {
		go consumer.Consume()
	}

	log.Printf("%d worker(s) of type %s started", len(consumers), workerType)
	<-forever

	return nil
}

func (s *Service) getConsumers(workerType lib.WorkerType) ([]IConsumer, error) {
	switch workerType {
	case lib.All:
		return s.getAllConsumers(), nil
	case lib.Book:
		return []IConsumer{s.BookConsumer}, nil
	default:
		return nil, errors.ErrInvalidWorkerType
	}
}

func (s *Service) getAllConsumers() []IConsumer {
	return []IConsumer{
		s.BookConsumer,
	}
}
