package usecase

import (
	"context"

	"github.com/mizuochikeita/homeway/client/domain/tcp"
)

type HomewayServerClient interface {
	OpenProxy(ctx context.Context) (chan tcp.Operation, error)
}
