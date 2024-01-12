package server

import "github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"

type PreFetcher struct {
	coin lib.Coin
}

func NewPreFetcher(cfg lib.Config) *PreFetcher {
	p := PreFetcher{
		coin: cfg.GetCoin(),
	}
	return &p
}

type BlockProcessor struct {
}
