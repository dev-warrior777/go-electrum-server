package server

import (
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
	GetMempoolHashes() ([]string, error)
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
