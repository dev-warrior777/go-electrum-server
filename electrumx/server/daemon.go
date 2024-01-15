package server

import (
	"context"
	"errors"

	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
)

// Daemon defines most behaviors of a btc-clone RPC node.
//
// Known daemons:
// - BtcDaemon
type Daemon interface {
	GetBlockCount() int64
	GetBlockHash(height int64) (string, error)
	GetBlockHashes(height int64, count int) ([]string, int, error)
}

func DaemonForCoin(coin lib.Coin) (Daemon, error) {
	switch coin.(type) {
	case lib.Bitcoin:
		daemonUrl := coin.GetDaemonUrl()
		rpcClient := NewDaemonClient(
			daemonUrl.Endpoint, daemonUrl.Username, daemonUrl.Password)
		daemon := BtcDaemon{
			daemonUrl: daemonUrl,
			rpcClient: rpcClient,
		}
		return daemon, nil
	}
	return nil, errors.New("unknown coin")
}

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
