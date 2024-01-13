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
	Tip() int64
	GetBlockHash(height int64) (string, error)
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

func (b BtcDaemon) Tip() int64 {

	return 0
}

func (b BtcDaemon) GetBlockHash(height int64) (string, error) {

	return "", nil
}
