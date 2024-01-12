package server

import "github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"

type Controller struct {
}

func NewController(cfg lib.Config) *Controller {
	c := Controller{}
	return &c
}
