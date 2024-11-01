package model

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Block struct {
	Number       *big.Int
	Hash         common.Hash
	ParentHash   common.Hash
	Timestamp    uint64
	Transactions []common.Hash
}
