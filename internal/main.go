package internal

import (
	"broker/internal/storages"
	"os"
)

func Main() {
	var store storages.PostgresStorage

	store.Init()

	ss := SocketServer{Storage: &store, LocalAddress: os.Args[1]}
	ss.RunServer()
}
