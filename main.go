package main

// https://pkg.go.dev/github.com/btcsuite/goleveldb

import (
	"context"
	"os"
	"os/signal"

	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
	"github.com/dev-warrior777/go-electrum-server.git/electrumx/server"
	"go.uber.org/zap"
)

func init() {
	// set up a global threadsafe zap logger
	logger := zap.Must(zap.NewDevelopment())
	zap.ReplaceGlobals(logger)
}

func main() {
	zap.S().Infof("ElectrumX: version: %s protocol %s", lib.Version(), lib.Protocol())

	// set up config - bitcoin regtest only for now
	cfg := lib.GetDefaultConfig()
	bitcoin, err := lib.NewBitcoin(lib.Regtest)
	if err != nil {
		zap.S().Errorf("%v\n", err)
		os.Exit(1)
	}
	cfg.SetCoin(bitcoin)

	// start controller
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	ctl := server.NewController(cfg)
	err = ctl.StartServer(ctx)
	if err != nil {
		zap.S().Errorf("%v\n", err)
		os.Exit(1)
	}

	<-ctx.Done()
	ctl.StopServer()
	cancel()

	zap.L().Info("exit")
}
