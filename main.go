package main

import (
	"broker/config"
	socketserver "broker/socket_server"
	"broker/storage"
)

func main() {
	conf := config.LoadConfig()

	var store storage.Storage
	store.Init()
	defer store.Close() // ?
	/*go ws_server.RunWS(store, ":"+conf.WebsoketPort)*/

	ss := socketserver.SocketServer{}
	ss.Init(store, conf.Address, conf.Type)
	ss.RunServer()
}
