package dao

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nocturna-ta/blockchain/internal/domain/model"
	"github.com/nocturna-ta/blockchain/internal/domain/repository"
	"github.com/nocturna-ta/blockchain/pkg/contracts"
	"github.com/nocturna-ta/golib/tracing"
	"math/big"
)

type BlockchainRepository struct {
	client    *ethclient.Client
	contracts *contracts.SimpleStorage
}

type OptsBlockchainRepository struct {
	Client          *ethclient.Client
	ContractAddress common.Address
}

func NewBlockchainRepository(opts *OptsBlockchainRepository) (repository.BlockchainRepository, error) {
	contract, err := contracts.NewSimpleStorage(opts.ContractAddress, opts.Client)
	if err != nil {
		return nil, err
	}
	return &BlockchainRepository{
		client:    opts.Client,
		contracts: contract,
	}, nil
}

func (b *BlockchainRepository) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "BlockchainRepository.GetBalance")
	defer span.End()
	return b.client.BalanceAt(ctx, address, nil)
}

func (b *BlockchainRepository) GetTransaction(ctx context.Context, hash common.Hash) (*model.Transaction, bool, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "BlockchainRepository.GetTransaction")
	defer span.End()

	tx, isPending, err := b.client.TransactionByHash(ctx, hash)
	if err != nil {
		return nil, false, err
	}

	var from common.Address
	if isPending {
		from, err = b.client.TransactionSender(ctx, tx, tx.Hash(), 0)
	} else {
		receipt, err := b.client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, false, err
		}
		from, err = b.client.TransactionSender(ctx, tx, receipt.BlockHash, receipt.TransactionIndex)
	}
	if err != nil {
		return nil, false, err
	}

	var to common.Address
	if tx.To() != nil {
		to = *tx.To()
	}

	return &model.Transaction{
		Hash:     tx.Hash(),
		From:     from,
		To:       to,
		Value:    tx.Value(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Nonce:    tx.Nonce(),
	}, isPending, nil

}

func (b *BlockchainRepository) GetBlock(ctx context.Context, number *big.Int) (*model.Block, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "BlockchainRepository.GetBlock")
	defer span.End()

	block, err := b.client.BlockByNumber(ctx, number)
	if err != nil {
		return nil, err
	}

	txHashes := make([]common.Hash, len(block.Transactions()))
	for i, tx := range block.Transactions() {
		txHashes[i] = tx.Hash()
	}

	return &model.Block{
		Number:       block.Number(),
		Hash:         block.Hash(),
		ParentHash:   block.ParentHash(),
		Timestamp:    block.Time(),
		Transactions: txHashes,
	}, nil
}

func (b *BlockchainRepository) GetContractValue(ctx context.Context) (*big.Int, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "BlockchainRepository.GetContractValue")
	defer span.End()

	return b.contracts.Get(nil)
}

func (b *BlockchainRepository) SetContractValue(ctx context.Context, privateKey *ecdsa.PrivateKey, from common.Address, value *big.Int) (common.Hash, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "BlockchainRepository.SetContractValue")
	defer span.End()

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return common.Hash{}, err
	}

	auth.Context = ctx

	tx, err := b.contracts.Set(auth, value)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}
