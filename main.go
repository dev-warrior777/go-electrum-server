package main

// https://pkg.go.dev/github.com/btcsuite/goleveldb#section-readme
import (
	"fmt"
	"os"

	"github.com/btcsuite/goleveldb/leveldb"
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

	// set up config - bitcoin only for now
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

func opendb() {
	db, err := leveldb.OpenFile("dbdata", nil)
	if err != nil {
		zap.S().Errorf("%v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	fmt.Println("opened db")

	err = db.Put([]byte("1"), []byte("one"), nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	b, err := db.Get([]byte("1"), nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("1=%v\n", string(b))
	err = db.Delete([]byte("1"), nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
