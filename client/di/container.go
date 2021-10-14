package di

import (
	"github.com/mizuochikeita/homeway/client/infra/environment"
	"github.com/mizuochikeita/homeway/client/usecase"
)

type Container struct {
	ConfigCache *environment.Config
}

func (c *Container) Config() *environment.Config {
	if c.ConfigCache != nil {
		return c.ConfigCache
	}
	conf, err := environment.ReadFromEnv()
	if err != nil {
		panic(err)
	}
	c.ConfigCache = conf
	return c.Config()
}

func (c *Container) ListenViaProxy() usecase.ListenViaProxy {
	return usecase.NewListenViaProxy(c.Client())
}

func (c *Container) Client() usecase.Client {
	return nil
}
