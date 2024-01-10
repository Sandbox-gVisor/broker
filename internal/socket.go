package internal

import (
	"broker/internal/storages"
	"encoding/json"
	"io"
	"log"
	"net"
)

type SocketServer struct {
	Storage      storages.Storage
	LocalAddress string
}

// RunServer runs server with config and waits for connections.
// Server will listen on address that is written in serv.LocalAddress
func (serv *SocketServer) RunServer() {
	log.Println("Server Running...")

	laddr, err := net.ResolveTCPAddr("tcp", serv.LocalAddress)
	if err != nil {
		log.Fatal(err.Error())
	}

	server, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer server.Close()
	log.Println("Listening on " + laddr.String())
	for {
		log.Println("Waiting for client...")
		connection, err := server.Accept()
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("Client connected " + connection.LocalAddr().String())
		go serv.ProcessClient(connection)
	}
}

// ProcessClient handles connected client
func (serv *SocketServer) ProcessClient(connection net.Conn) {
	dec := json.NewDecoder(connection)

	for {
		var logs map[string]interface{}
		err := dec.Decode(&logs)

		if err != nil {
			log.Println("Error reading: ", err.Error())
			if err == io.EOF {
				break
			}

			continue
		}

		jsonLogs, err := json.Marshal(&logs)
		if err != nil {
			log.Println("Error while marshaling logs: ", err.Error())
			continue
		}

		serv.Storage.SaveMessage(string(jsonLogs))
	}

	log.Println("Closing connection with client: " + connection.LocalAddr().String() + "...")
	connection.Close()

	serv.Storage.FlushStorage()
}
