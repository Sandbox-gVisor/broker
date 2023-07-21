package rabbit

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func getQueueName(ch *amqp.Channel, queueName string) string {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Panic(err)
	}
	return q.Name
}

func getConsume(ch *amqp.Channel, queueName string) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		getQueueName(ch, queueName),
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Panic(err)
	}
	return msgs
}

func publishMessage(ch *amqp.Channel, ctx context.Context, queueName string, body string) {
	err := ch.PublishWithContext(ctx,
		"",                          // exchange
		getQueueName(ch, queueName), // routing key
		false,                       // mandatory
		false,                       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Panic(err)
	}
}
