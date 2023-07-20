package socketserver

import (
	"fmt"
	"log"
	"net"

	"broker/config"
)

type SocketServer struct {
	Address string
	Type    string
}

func (self *SocketServer) Init(config config.Config) {
	self.Address = config.Address
	self.Type = config.Type
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
	if err != nil {
		log.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
	connection.Close()
}
