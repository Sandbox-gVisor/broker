package main

import (
	socketserver "broker/socket_server"
	"broker/storage"
	"log"
	"os"
)

func main() {
	var store storage.Storage
	store.Init()
	defer store.Close()
	log.Println(store.RedisClient.String())

	ss := socketserver.SocketServer{Broker: store, LocalAddress: os.Args[1]}
	ss.RunServer()
}
