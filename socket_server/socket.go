package socketserver

import (
	"broker/storage"
	"encoding/json"
	"fmt"
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
	//defer server.Close()
	log.Print("Listening on " + serv.Address)
	for {
		connection, err := server.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("client connected")
		go serv.ProcessClient(connection)
	}
}

// ProcessClient handles connected client
func (serv *SocketServer) ProcessClient(connection net.Conn) {
	dec := json.NewDecoder(connection)

	for {
		var logs map[string]interface{}
		err := dec.Decode(&logs)

		fmt.Println(logs)

		if err != nil {
			log.Println("Error reading:", err.Error())
			continue
			//return
		}

		jsonLogs, err := json.Marshal(&logs)
		if err != nil {
			log.Println("Error while marshaling logs:", err.Error())
			continue
			//return
		}

		serv.Broker.AddString(string(jsonLogs))
	}
	/*buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	defer connection.Close()
	if err != nil {
		log.Println("Error reading:", err.Error())
	}
	data := string(buffer[:mLen])
	serv.Broker.AddString(data)
	fmt.Println(data)*/
}
