package lib

import (
	"errors"

	"github.com/btcsuite/btcd/chaincfg"
)

type Net string

var (
	Regtest Net = "regtest"
	Testnet Net = "testnet"
	Mainnet Net = "regtest"
)

type Coin interface {
	GetDaemonUrl() *DaemonUrl
	//TODO: more methods
}

type Bitcoin struct {
	Params    *chaincfg.Params
	DaemonUrl *DaemonUrl
}

type DaemonUrl struct {
	Endpoint string
	Username string
	Password string
}

var (
	// harness alpha
	defaultBitcoinRegtestUrl = DaemonUrl{
		Endpoint: "http://127.0.0.1:20556",
		Username: "user",
		Password: "pass",
	}
	defaultBitcoinTestnetUrl = DaemonUrl{
		Endpoint: "http://127.0.0.1:18332",
		Username: "???",
		Password: "???",
	}
	defaultBitcoinMainnetUrl = DaemonUrl{
		Endpoint: "http://127.0.0.1:8332",
		Username: "???",
		Password: "???",
	}
)

func NewBitcoin(net Net) (Coin, error) {
	var params *chaincfg.Params
	var daemonUrl *DaemonUrl
	switch net {
	case Regtest:
		params = &chaincfg.RegressionNetParams
		daemonUrl = &defaultBitcoinRegtestUrl
	case Testnet:
		params = &chaincfg.TestNet3Params
		daemonUrl = &defaultBitcoinTestnetUrl
	case Mainnet:
		params = &chaincfg.MainNetParams
		daemonUrl = &defaultBitcoinMainnetUrl
	default:
		return nil, errors.New("invalid network")
	}
	bitcoin := Bitcoin{
		Params:    params,
		DaemonUrl: daemonUrl,
	}
	return bitcoin, nil
}

func (b Bitcoin) GetDaemonUrl() *DaemonUrl {
	return b.DaemonUrl
}
