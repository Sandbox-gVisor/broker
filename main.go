package main

import (
	"fmt"
	"strconv"

	"broker/config"
	"broker/rabbit"
	socketserver "broker/socket_server"
	ws_server "broker/ws_server"
)

func main() {
	var mb rabbit.MessageBroker
	mb.Open("amqp://guest:guest@localhost:5672", "main")
	defer mb.Close()

	for i := 0; i < 10; i++ {
		mb.Send(strconv.Itoa(i))
	}
	messages := mb.Read()
	go func() {
		for m := range messages {
			fmt.Println(string(m.Body))
		}
	}()

	conf := config.LoadConfig()
	fmt.Println(conf)
	ss := socketserver.SocketServer{}
	ss.Init(conf)
	go ws_server.RunWS()
	ss.RunServer()
}
