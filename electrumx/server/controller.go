package server

import (
	"context"
	"errors"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
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

// StartServer initializes and starts the server
func (c *Controller) StartServer() error {
	zap.S().Info("starting server")
	daemon, err := DaemonForCoin(c.config.GetCoin())
	if err != nil {
		return err
	}
	c.daemon = daemon
	// check daemon genesis
	genesisBlockHash, _ := c.daemon.GetBlockHash(0)
	zap.S().Infof("daemon genesis blockhash: %s", genesisBlockHash)
	params := c.config.GetCoin().GetParams()
	trueGenesis := params.GenesisHash
	daemonGenesis, _ := chainhash.NewHashFromStr(genesisBlockHash)
	if !trueGenesis.IsEqual(daemonGenesis) {
		zap.S().Errorf("daemon genesis: %s", c.daemon.GetBlockCount())
		return errors.New("invalid daemon genesis block hash")
	}
	zap.S().Infof("daemon has valid genesis for net %s", params.Name)
	// daemon height
	blocks := c.daemon.GetBlockCount()
	zap.S().Infof("daemon blocks: %d", blocks)

	return nil
}

// StopServer gracefully stops the server
func (c *Controller) StopServer() {
	zap.S().Info("stopping server")

}
