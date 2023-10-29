package svc

import (
	"composeapp/ent"
	"composeapp/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Client *ent.Client
}

func NewServiceContext(c config.Config, con *ent.Client) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Client: con,
	}
}
