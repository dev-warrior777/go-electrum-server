package server

import (
	"context"

	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
	"go.uber.org/zap"
)

type Controller struct {
	ctx    context.Context
	config *lib.Config
	daemon Daemon
}

func NewController(ctx context.Context, cfg *lib.Config) *Controller {
	c := Controller{
		ctx:    ctx,
		config: cfg,
	}
	zap.S().Info("new controller")
	return &c
}

// StopServer initailaizes and starts the server
func (c *Controller) StartServer() error {
	zap.S().Info("starting server")
	daemon, err := DaemonForCoin(c.config.GetCoin())
	if err != nil {
		return err
	}
	c.daemon = daemon
	zap.S().Infof("daemon tip: %d", c.daemon.Tip())

	genesisBlockHash, _ := c.daemon.GetBlockHash(0)
	zap.S().Infof("daemon genesis blockhash: %s", genesisBlockHash)

	return nil
}

// StopServer gracefully stops the server
func (c *Controller) StopServer() {
	zap.S().Info("stopping server")

}
