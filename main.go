package main

import (
	"broker/config"
	"broker/rabbit"
	socketserver "broker/socket_server"
	ws_server "broker/ws_server"
)

func main() {
	conf := config.LoadConfig()

	var mb rabbit.MessageBroker
	mb.Open(conf.RabbitAddress, conf.QueueName)
	defer mb.Close()
	go ws_server.RunWS(mb, ":"+conf.WebsoketPort)

	ss := socketserver.SocketServer{}
	ss.Init(mb, conf.Address, conf.Type)
	ss.RunServer()
}
