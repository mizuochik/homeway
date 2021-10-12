package di

import "github.com/mizuochikeita/homeway/client/infra/environment"

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
