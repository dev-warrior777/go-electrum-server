package main

// https://pkg.go.dev/github.com/btcsuite/goleveldb

import (
	"os"

	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
	"go.uber.org/zap"
)

func init() {
	// set up a global threadsafe zap logger
	logger := zap.Must(zap.NewDevelopment())
	zap.ReplaceGlobals(logger)
}

func main() {
	zap.L().Info("ElectrumX", zap.String("version", lib.Version()))

	// set up config - bitcoin regtest only for now
	cfg := lib.GetDefaultConfig()
	bitcoin, err := lib.NewBitcoin(lib.Regtest)
	if err != nil {
		zap.S().Errorf("%v\n", err)
		os.Exit(1)
	}
	cfg.SetCoin(bitcoin)

	// start controller

	zap.L().Info("exit")
}
