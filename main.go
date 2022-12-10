package main

import (
	"log"
	"time"

	"mq/consumer"
	"mq/lib"

	bookConsumer "mq/consumer/book"
	bookProducer "mq/producer/book"
)

func main() {
	shotDown := make(chan bool)

	// Consumer
	mqConnection := lib.NewConnection()
	bookConsumer_ := bookConsumer.NewConsumer(mqConnection)
	consumerService := consumer.NewConsumeService(bookConsumer_)

	go func() {
		err := consumerService.Consume(lib.Book)
		if err != nil {
			log.Println("Consumer could not run")
		}
	}()

	// Producer
	bookProducer_ := bookProducer.NewProducer(mqConnection)
	for _, i := range []string{"1", "2", "3", "4", "5", "6"} {
		err := bookProducer_.Send(i, "Burkay")
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Message send")
		}

		durationTime := time.Second * 4
		time.Sleep(durationTime)
	}

	<-shotDown
}
