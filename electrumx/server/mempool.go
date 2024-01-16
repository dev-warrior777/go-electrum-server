package server

import (
	"context"

	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
)

type Mempool struct {
}

func NewMempool(cfg *lib.Config) *Mempool {
	mp := Mempool{}
	return &mp
}

func (mp *Mempool) run(ctx context.Context) error {

	return nil
}
