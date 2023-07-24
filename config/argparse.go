package config

import (
	"os"

	"github.com/akamensky/argparse"
)

func parseCliArgs() (Config, error) {
	parser := argparse.NewParser("sandbox-broker", "Tool for connect gVisor and web interface")
	address := parser.String("a", "address", &argparse.Options{Required: true, Help: "Socket address"})
	socketType := parser.String("t", "type", &argparse.Options{Required: true, Help: "Socket type"})
	rabbit := parser.String("r", "rabbit", &argparse.Options{Required: true, Help: "Rabbit address"})
	queue := parser.String("q", "queue", &argparse.Options{Required: false, Help: "Rabbit queue"})
	port := parser.String("p", "port", &argparse.Options{Required: true, Help: "WebSoket port"})

	err := parser.Parse(os.Args)
	if err != nil {
		return Config{}, err
	}
	q := *queue
	if q == "" {
		q = "main"
	}
	return Config{
		Address:       *address,
		Type:          *socketType,
		WebsoketPort:  *port,
		RabbitAddress: *rabbit,
		QueueName:     q,
	}, nil
}
