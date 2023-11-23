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
	if err := store.RedisClient.Ping(store.Ctx).Err(); err != nil {
		log.Fatal("Couldn't connect to Redis database: ", err)
		return
	} else {
		log.Println("Redis Ping was successful!")
	}
	defer store.Close()

	log.Println(store.RedisClient.String())

	ss := socketserver.SocketServer{Broker: store, LocalAddress: os.Args[1]}
	ss.RunServer()
}
