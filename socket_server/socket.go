package socketserver

import (
	"fmt"
	"log"
	"net"

	"broker/rabbit"
)

type SocketServer struct {
	Broker  rabbit.MessageBroker
	Address string
	Type    string
}

func (self *SocketServer) Init(broker rabbit.MessageBroker, address string, t string) {
	self.Broker = broker
	self.Address = address
	self.Type = t
}

func (self *SocketServer) RunServer() {
	fmt.Println("Server Running...")
	server, err := net.Listen(self.Type, self.Address)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer server.Close()
	log.Print("Listening on " + self.Address)
	for {
		connection, err := server.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("client connected")
		go self.ProcessClient(connection)
	}

}

func (self *SocketServer) ProcessClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	defer connection.Close()
	if err != nil {
		log.Println("Error reading:", err.Error())
	}
	data := string(buffer[:mLen])
	self.Broker.Send(data)
}
