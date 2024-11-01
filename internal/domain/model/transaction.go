package model

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Transaction struct {
	Hash     common.Hash
	From     common.Address
	To       common.Address
	Value    *big.Int
	Gas      uint64
	GasPrice *big.Int
	Nonce    uint64
}
