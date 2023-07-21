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
	mb.Open("amqp://guest:guest@localhost:5672")
	defer mb.Close()

	for i := 0; i < 10; i++ {
		mb.SendToQueue("main", strconv.Itoa(i))
	}
	messages := mb.Read("main")
	fmt.Println(messages)
	for i, m := range messages {
		fmt.Println(i, m)
	}

	conf := config.LoadConfig()
	fmt.Println(conf)
	ss := socketserver.SocketServer{}
	ss.Init(conf)
	go ws_server.RunWS()
	ss.RunServer()
}
