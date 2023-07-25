package socketserver

import (
	"fmt"
	"log"
	"net"

	"broker/storage"
)

type SocketServer struct {
	Broker  storage.Storage
	Address string
	Type    string
}

func (self *SocketServer) Init(broker storage.Storage, address string, t string) {
	self.Broker = broker
	self.Address = address
	self.Type = t
}

func (serv *SocketServer) RunServer() {
	fmt.Println("Server Running...")
	server, err := net.Listen(serv.Type, serv.Address)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer server.Close()
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

func (serv *SocketServer) ProcessClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	defer connection.Close()
	if err != nil {
		log.Println("Error reading:", err.Error())
	}
	data := string(buffer[:mLen])
	serv.Broker.Send(data)
	fmt.Println(data)
}
