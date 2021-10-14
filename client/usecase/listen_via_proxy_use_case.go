package usecase

import (
	"context"
)

type ListenViaProxy interface {
	Run(ctx *context.Context) error
}

type listenViaProxy struct {
	client HomewayServerClient
}

func NewListenViaProxy(client HomewayServerClient) ListenViaProxy {
	return &listenViaProxy{
		client: client,
	}
}

func (uc *listenViaProxy) Run(ctx *context.Context) error {
	return nil
}
