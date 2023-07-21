package main

import (
	"fmt"

	"broker/config"
	socketserver "broker/socket_server"
	ws_server "broker/ws_server"
)

func main() {
	conf := config.LoadConfig()
	fmt.Println(conf)
	ss := socketserver.SocketServer{}
	ss.Init(conf)
	go ws_server.RunWS()
	ss.RunServer()
}
