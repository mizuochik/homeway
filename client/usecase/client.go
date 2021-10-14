package usecase

import (
	"context"

	"github.com/mizuochikeita/homeway/client/domain/tcp"
)

type Client interface {
	OpenProxy(ctx context.Context) (chan tcp.Operation, error)
}
