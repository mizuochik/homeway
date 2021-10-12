package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/mizuochikeita/homeway/client/di"
)

func main() {
	dc := di.Container{}
	endpoint := dc.Config().ServerURL.ResolveReference(&url.URL{Path: "/proxy"})
	conn, res, err := websocket.DefaultDialer.Dial(endpoint.String(), nil)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("error: %s", err)
			log.Printf("< %s", string(msg))
		}
	}
}
