package server

import (
	"net"
	"net/http"

	"github.com/mizuochikeita/homeway/infra/http/server"
)

type Container struct {
}

func (c *Container) Server() *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/proxy", c.ProxyHandler())
	return &http.Server{
		Addr:    net.JoinHostPort("localhost", "8080"),
		Handler: mux,
	}
}

func (*Container) ProxyHandler() http.Handler {
	return server.ProxyHandler()
}
