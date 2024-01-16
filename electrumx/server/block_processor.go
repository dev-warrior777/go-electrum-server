package server

import (
	"context"
	"sync"

	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
)

const (
	MIN_CACHE_SIZE = 10 * 1024 * 1024
)

type PreFetchCache struct {
	size        uint32
	averageSize uint32
}

type PreFetcher struct {
	daemon            Daemon
	caughtUp          bool
	fetchedHeight     int64
	fetchedHeightsMtx sync.Mutex
	pollingDelaySecs  int
	cache             *PreFetchCache
}

func NewPreFetcher(cfg *lib.Config, daemon Daemon) *PreFetcher {
	pf := PreFetcher{
		daemon:           daemon,
		pollingDelaySecs: cfg.PreFetchPollingDelaySecs,
		cache: &PreFetchCache{
			size:        MIN_CACHE_SIZE,
			averageSize: MIN_CACHE_SIZE,
		},
	}
	return &pf
}

// main pre-fetch loop
func (pf *PreFetcher) Run(ctx context.Context) error {

	return nil
}

type BlockProcessor struct {
	preFetcher *PreFetcher
	db         *DB
}

func NewBlockProcessor(cfg *lib.Config, daemon Daemon, db *DB) *BlockProcessor {
	bp := BlockProcessor{
		preFetcher: NewPreFetcher(cfg, daemon),
		db:         db,
	}
	return &bp
}

// main block-processor loop fetches, process & index pre-fetched blocks
func (bp *BlockProcessor) run(ctx context.Context) error {

	return nil
}
