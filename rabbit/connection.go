package rabbit

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageBroker struct {
	Connection *amqp.Connection
}

func (mb *MessageBroker) Open(address string) {
	var err error
	mb.Connection, err = amqp.Dial(address)
	if err != nil {
		log.Fatalf("Can't connect to RabbitMQ server: %s", err)
	}
}

func (mb *MessageBroker) Close() {
	err := mb.Connection.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (mb *MessageBroker) SendToQueue(queueName string, body string) {
	ch, err := mb.Connection.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel. Error: %s", err)
	}

	defer func() {
		_ = ch.Close() // Закрываем подключение в случае удачной попытки подключения
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	publishMessage(ch, ctx, queueName, body)
	defer cancel()
	log.Printf(" [x] Sent %s\n", body)
}

func (mb *MessageBroker) Read(queueName string) []string {
	ch, err := mb.Connection.Channel()
	if err != nil {
		log.Panic(err)
	}
	defer ch.Close()

	var (
		result  []string
		forever chan struct{}
	)

	msgs := getConsume(ch, queueName)
	fmt.Println(msgs)
	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
			result = append(result, string(d.Body))
		}
	}()
	fmt.Println("here")
	<-forever
	return result
}
