package ws_server

import (
	"broker/storage"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const (
	PullCmd = "pull"
)

func RunWS(broker storage.Storage, port string) {
	fmt.Println("Running ws")

	err := http.ListenAndServe(port,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			go func() {
				var conn, _, _, err = ws.UpgradeHTTP(r, w)
				if err != nil {
					log.Fatal(err)
				defer conn.Close()

				for {
					msg, _, err := wsutil.ReadClientData(conn)
					fmt.Println(string(msg))
					if err != nil {
						continue
					}
					/*
						if string(msg) == PullCmd { // if pull from client, subscrube to broker
							messages := broker.Read()
							go func() {
								for m := range messages {
									err = wsutil.WriteServerMessage(conn, op, m.Body)
									if err != nil {
										log.Println(err)
										continue
									}
								}
							}()
						}
					*/
				}

				handleConnection(conn, broker)
			}()
		}))
	if err != nil {
		return
	}
}

func handleConnection(conn net.Conn, broker rabbit.MessageBroker) {
	defer conn.Close()

	for {
		msg, op, err := wsutil.ReadClientData(conn)
		fmt.Println(string(msg))
		if err != nil {
			log.Println(err)
			continue
		}

		if string(msg) == PullCmd { // if pull from client, subscribe to broker
			messages := broker.Read()
			writeToSocket(messages, conn, op)
		}
	}
}

func writeToSocket(messages <-chan amqp.Delivery, conn net.Conn, op ws.OpCode) {
	for m := range messages {
		err := wsutil.WriteServerMessage(conn, op, m.Body)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
