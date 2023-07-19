package main

import (
	"fmt"

	"broker/config"
	socketserver "broker/socket_server"
)

func main() {
	conf := config.LoadConfig()
	fmt.Println(conf)
	ss := socketserver.SocketServer{}
	ss.Init(conf)
	ss.RunServer()
}
