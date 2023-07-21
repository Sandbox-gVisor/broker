package main

import (
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
	// TODO delete it
	for i := 0; i < 10; i++ {
		mb.Send(strconv.Itoa(i))
	}
	// ---
	conf := config.LoadConfig()
	go ws_server.RunWS(mb)
	ss := socketserver.SocketServer{}
	ss.Init(conf, mb)
	ss.RunServer()
}
