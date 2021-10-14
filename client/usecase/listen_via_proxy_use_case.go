package usecase

import (
	"context"
)

type ListenViaProxy interface {
	Run(ctx *context.Context) error
}

type listenViaProxy struct {
}

func NewListenViaProxy() ListenViaProxy {
	return &listenViaProxy{}
}

func (uc *listenViaProxy) Run(ctx *context.Context) error {
	return nil
}
