package blockchain

import (
	"github.com/nocturna-ta/blockchain/internal/domain/repository"
	"github.com/nocturna-ta/blockchain/internal/usecases"
)

type Module struct {
	blockchainRepo repository.BlockchainRepository
}

type Opts struct {
	BlockchainRepo repository.BlockchainRepository
}

func New(opts *Opts) usecases.BlockchainUseCases {
	return &Module{
		blockchainRepo: opts.BlockchainRepo,
	}
}
