package config

import (
	"os"

	"github.com/akamensky/argparse"
)

func parseCliArgs() (Config, error) {
	parser := argparse.NewParser("sandbox-cli", "tool for in-time configuraion gVisor")
	address := parser.String("a", "address", &argparse.Options{Required: true, Help: "Socket address"})
	socketType := parser.String("t", "type", &argparse.Options{Required: true, Help: "Socket type"})

	err := parser.Parse(os.Args)
	if err != nil {
		return Config{}, err
	}
	return Config{Address: *address, Type: *socketType}, nil
}
