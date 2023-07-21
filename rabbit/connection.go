package rabbit

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	contextLifeTime = 5
)

type MessageBroker struct {
	QueueName  string
	Connection *amqp.Connection
	ReadChan   *amqp.Channel
	WriteChan  *amqp.Channel
}

func (mb *MessageBroker) Open(address string, queueName string) {
	var (
		err error
		ch  *amqp.Channel
	)
	mb.QueueName = queueName
	mb.Connection, err = amqp.Dial(address)
	if err != nil {
		log.Fatalf("Can't connect to RabbitMQ server: %s", err)
	}
	ch, err = mb.Connection.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel. Error: %s", err)
	}
	mb.WriteChan = ch
	ch, err = mb.Connection.Channel()
	if err != nil {
		log.Panic(err)
	}
	mb.ReadChan = ch
}

func (mb *MessageBroker) Close() {
	var err error
	err = mb.ReadChan.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = mb.WriteChan.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = mb.Connection.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (mb *MessageBroker) Send(body string) {
	ctx, cancel := context.WithTimeout(context.Background(), contextLifeTime*time.Second)
	defer cancel()

	publishMessage(mb.WriteChan, ctx, mb.QueueName, body)

	log.Printf("Sent %s\n", body)
}

func (mb *MessageBroker) Read() <-chan amqp.Delivery {
	return getConsume(mb.ReadChan, mb.QueueName)
}
