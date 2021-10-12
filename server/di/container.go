package di

import (
	"net"
	"net/http"

	hshttp "github.com/mizuochikeita/homeway/server/infra/http"
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
	return hshttp.ProxyHandler()
}
