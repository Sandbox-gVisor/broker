package ws_server

import (
	"broker/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const (
	PullCmd = "pull"
)

func RunWS(broker storage.Storage, port string) {
	fmt.Println("Running ws")
	http.ListenAndServe(port,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			conn, _, _, err := ws.UpgradeHTTP(r, w)
			if err != nil {
				log.Fatal(err)
			}
			go func() {
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
			}()
		}))
}
