package server

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nocturna-ta/blockchain/config"
	"github.com/nocturna-ta/blockchain/internal/interfaces/dao"
	"github.com/nocturna-ta/blockchain/internal/usecases"
	"github.com/nocturna-ta/blockchain/internal/usecases/blockchain"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/txmanager"
	txSql "github.com/nocturna-ta/golib/txmanager/sql"
)

type container struct {
	Cfg          config.MainConfig
	BlockchainUc usecases.BlockchainUseCases
}

type options struct {
	Cfg    *config.MainConfig
	DB     *sql.Store
	Client *ethclient.Client
}

func newContainer(opts *options) *container {
	blockchainRepo, err := dao.NewBlockchainRepository(&dao.OptsBlockchainRepository{
		Client:          opts.Client,
		ContractAddress: common.HexToAddress(opts.Cfg.Blockchain.ContractAddress),
	})
	if err != nil {
		log.Fatal("Failed to initiate blockchain repository")
	}

	_, err = txmanager.New(context.Background(), &txmanager.DriverConfig{
		Type: "sql",
		Config: txSql.Config{
			DB: opts.DB,
		},
	})
	if err != nil {
		log.Fatal("Failed to instantiate transaction manager ")
	}

	blockchainUc := blockchain.New(&blockchain.Opts{
		BlockchainRepo: blockchainRepo,
	})
	return &container{
		Cfg:          *opts.Cfg,
		BlockchainUc: blockchainUc,
	}
}
