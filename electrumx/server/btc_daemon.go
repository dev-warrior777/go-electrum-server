package server

import (
	"context"

	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
)

type BtcDaemon struct {
	daemonUrl *lib.DaemonUrl // no failover urls for now
	rpcClient *DaemonClient
}

func (b BtcDaemon) GetBlockCount() int64 {
	i, err := b.rpcClient.GetBlockCount(context.Background())
	if err != nil {
		return -1
	}
	return int64(i)
}

func (b BtcDaemon) GetBlockHash(height int64) (string, error) {
	hash, err := b.rpcClient.GetBlockHash(context.Background(), height)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (b BtcDaemon) GetBlockHashes(height int64, count int) ([]string, int, error) {
	hashes, i, err := b.rpcClient.GetBlockHashes(context.Background(), height, count)
	if err != nil {
		return nil, i, err
	}
	return hashes, i, nil
}

func (b BtcDaemon) GetMempoolHashes() ([]string, error) {
	hashes, err := b.rpcClient.GetMempoolHashes(context.Background())
	if err != nil {
		return nil, err
	}
	return hashes, nil
}
