package server

import (
	"context"
	"github.com/nocturna-ta/blockchain/config"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/txmanager"
	txSql "github.com/nocturna-ta/golib/txmanager/sql"
)

type container struct {
	Cfg config.MainConfig
}

type options struct {
	Cfg *config.MainConfig
	DB  *sql.Store
}

func newContainer(opts *options) *container {
	txMgr, err := txmanager.New(context.Background(), &txmanager.DriverConfig{
		Type: "sql",
		Config: txSql.Config{
			DB: opts.DB,
		},
	})
	if err != nil {
		log.Fatal("Failed to instantiate transaction manager ")
	}
	return &container{
		Cfg: *opts.Cfg,
	}
}
