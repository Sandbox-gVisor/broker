package socketserver

import (
	"broker/storage"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type SocketServer struct {
	Broker  storage.Storage
	Address string
	Type    string
}

type Log struct {
}

// Init configures server socket for running. Broker and address are taken from config
func (serv *SocketServer) Init(broker storage.Storage, address string, t string) {
	serv.Broker = broker
	serv.Address = address
	serv.Type = t
}

// RunServer runs server with config and waits for connections.
// Server will listen on address that is written in serv.Address
func (serv *SocketServer) RunServer() {
	fmt.Println("Server Running...")

	_ = os.Remove(serv.Address)
	server, err := net.Listen(serv.Type, serv.Address)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer server.Close()
	log.Print("Listening on " + serv.Address)
	for {
		log.Println("Waiting for client...")
		connection, err := server.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("Client connected!")
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
			log.Println("Error reading:", err.Error())
			if err == io.EOF {
				break
			}

			continue
		}

		jsonLogs, err := json.Marshal(&logs)
		if err != nil {
			log.Println("Error while marshaling logs:", err.Error())
			continue
		}

		serv.Broker.AddString(string(jsonLogs))
	}

	log.Println("Closing connection with client...")
	connection.Close()
}
