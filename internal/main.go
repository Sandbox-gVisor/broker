package internal

import (
	"broker/internal/storages"
	"log"
	"os"
)

func Main() {
	var store storages.RedisStorage

	store.Init()
	if err := store.RedisClient.Ping(store.Ctx).Err(); err != nil {
		log.Fatal("Couldn't connect to Redis database: ", err)
		return
	} else {
		log.Println("Redis Ping was successful!")
	}
	defer store.Close()

	log.Println(store.RedisClient.String())

	ss := SocketServer{Storage: &store, LocalAddress: os.Args[1]}
	ss.RunServer()
}
