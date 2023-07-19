package socketserver

import (
	"fmt"
	"log"
	"net"

	"broker/config"
)

type SocketServer struct {
	Host string
	Port string
	Type string
}

func (server *SocketServer) Init(config config.Config) {
	server.Host = config.Host
	server.Port = config.Port
	server.Type = config.Type
}

func (self *SocketServer) RunServer() {
	fmt.Println("Server Running...")
	address := self.Host + ":" + self.Port
	server, err := net.Listen(self.Type, address)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer server.Close()
	log.Print("Listening on " + address)
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
