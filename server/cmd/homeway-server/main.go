package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/mizuochikeita/homeway/server/di"
)

func main() {
	dc := &di.Container{}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	s := dc.Server()
	go func() {
		log.Printf("Listening on %s ...", s.Addr)
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("failed to listen and serve: %s", err)
		}
	}()
	<-ctx.Done()
	log.Println("Shutting down ...")
	if err := s.Shutdown(context.TODO()); err != nil {
		log.Printf("failed to shutdown: %s", err)
	}
}
