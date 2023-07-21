package ws_server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const (
	PullCmd = "pull"
)

func RunWS() {
	fmt.Println("Running ws")
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				fmt.Println(string(msg))
				if err != nil {
					log.Println(err)
					continue
				}
				if string(msg) == PullCmd {
					// TODO read data from message broker row by row
					err = wsutil.WriteServerMessage(conn, op, msg)
					if err != nil {
						log.Println(err)
						continue
					}
				}
			}
		}()
	}))
}
