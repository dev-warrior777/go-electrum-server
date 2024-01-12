package server

import (
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/dev-warrior777/go-electrum-server.git/electrumx/lib"
	"go.uber.org/zap"
)

type DB struct {
	db      *leveldb.DB
	datadir string
}

func NewDB(cfg lib.Config) *DB {
	db := DB{
		datadir: cfg.Datadir,
	}
	return &db
}

func (d *DB) Open() error {
	ldb, err := leveldb.OpenFile(db.datadir, nil)
	if err != nil {
		zap.S().Errorf("%v\n", err)
		return err
	}
	d.db = ldb
}
