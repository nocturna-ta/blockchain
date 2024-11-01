package repository

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nocturna-ta/blockchain/internal/domain/model"
	"math/big"
)

type BlockchainRepository interface {
	GetBalance(ctx context.Context, address common.Address) (*big.Int, error)
	GetTransaction(ctx context.Context, hash common.Hash) (*model.Transaction, bool, error)
	GetBlock(ctx context.Context, number *big.Int) (*model.Block, error)
	GetContractValue(ctx context.Context) (*big.Int, error)
	SetContractValue(ctx context.Context, privateKey *ecdsa.PrivateKey, from common.Address, value *big.Int) (common.Hash, error)
}
