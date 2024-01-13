package server

import (
	"context"

	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
	"go.uber.org/zap"
)

type Controller struct {
	context context.Context
	config  *lib.Config
}

func NewController(ctx context.Context, cfg *lib.Config) *Controller {
	c := Controller{
		context: ctx,
		config:  cfg,
	}
	zap.S().Info("new controller")
	return &c
}

// StopServer initailaizes and starts the server
func (c *Controller) StartServer() error {
	zap.S().Info("starting server")

	return nil
}

// StopServer gracefully stops the server
func (c *Controller) StopServer() {
	zap.S().Info("stopping server")

}
