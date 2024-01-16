package server

import (
	"context"
	"time"

	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
	"go.uber.org/zap"
)

const CONTROLLER_CHK_SECS = 5

// Block and Mempool send notifications on new blocks and new mempool hashes
type Notifications struct {
}

// Controller starts and coordinates the ElectrumX server
type Controller struct {
	config         *lib.Config
	daemon         Daemon
	blockProcessor *BlockProcessor
	mempool        *Mempool
	db             *DB
}

func NewController(cfg *lib.Config) *Controller {
	c := Controller{
		config: cfg,
	}
	return &c
}

// StartServer initializes and starts the server
func (c *Controller) StartServer(ctx context.Context) error {
	zap.S().Info("starting server")

	// make daemon client
	daemon, err := DaemonForCoin(c.config.GetCoin())
	if err != nil {
		return err
	}
	c.daemon = daemon

	// database
	c.db = NewDB(c.config)

	// block-processor
	c.blockProcessor = NewBlockProcessor(c.config, c.daemon, c.db)

	// mempool processor
	c.mempool = NewMempool(c.config)

	return c.run(ctx)
}

func (c *Controller) run(ctx context.Context) error {

	go func() {
		zap.S().Info("controller run")

		for {
			select {
			case <-ctx.Done():
				zap.S().Info("server shutdown in controller run")
				return
			case <-time.After(time.Second * CONTROLLER_CHK_SECS):
				zap.S().Info("controller run - tick")
			}
		}

	}()

	zap.S().Info("controller run exit")
	return nil
}

// StopServer gracefully stops the server
func (c *Controller) StopServer() {
	zap.S().Info("stopping server")

	zap.S().Info("server exit")
}
