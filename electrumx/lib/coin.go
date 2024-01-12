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
	//TODO: methods
}

type Bitcoin struct {
	Params chaincfg.Params
}

func NewBitcoin(net Net) (Coin, error) {
	var params chaincfg.Params
	switch net {
	case Regtest:
		params = chaincfg.RegressionNetParams
	case Testnet:
		params = chaincfg.TestNet3Params
	case Mainnet:
		params = chaincfg.MainNetParams
	default:
		return nil, errors.New("invalid network")
	}
	bitcoin := Bitcoin{
		Params: params,
	}
	return bitcoin, nil
}
