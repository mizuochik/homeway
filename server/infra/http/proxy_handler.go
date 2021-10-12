package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type DemoMessage struct {
	Body string `json:"body"`
}

func ProxyHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		conn, err := upgrader.Upgrade(rw, req, nil)
		if err != nil {
			return
		}
		if err := conn.WriteJSON(DemoMessage{Body: "hello homeway"}); err != nil {
			log.Printf("failed to write a message: %s", err)
		}
	})
}
