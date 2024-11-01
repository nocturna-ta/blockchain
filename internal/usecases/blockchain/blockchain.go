package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nocturna-ta/blockchain/internal/usecases/request"
	"github.com/nocturna-ta/blockchain/internal/usecases/response"
	"github.com/nocturna-ta/blockchain/pkg/helper"
	"github.com/nocturna-ta/golib/tracing"
	"math/big"
	"strconv"
)

func (m *Module) GetBalance(ctx context.Context, address string) (*response.BalanceResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "BlockchainUseCases.GetBalance")
	defer span.End()

	addr := common.HexToAddress(address)
	balance, err := m.blockchainRepo.GetBalance(ctx, addr)
	if err != nil {
		return nil, err
	}

	return &response.BalanceResponse{
		Address: addr.Hex(),
		Balance: balance.String(),
	}, nil
}

func (m *Module) GetTransaction(ctx context.Context, hash string) (*response.TransactionResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "BlockchainUseCases.GetTransaction")
	defer span.End()

	txHash := common.HexToHash(hash)
	tx, isPending, err := m.blockchainRepo.GetTransaction(ctx, txHash)
	if err != nil {
		return nil, err
	}

	return &response.TransactionResponse{
		Hash:      tx.Hash.Hex(),
		From:      tx.From.Hex(),
		To:        tx.To.Hex(),
		Value:     tx.Value.String(),
		Gas:       tx.Gas,
		GasPrice:  tx.GasPrice.String(),
		Nonce:     tx.Nonce,
		IsPending: isPending,
	}, nil
}

func (m *Module) GetBlock(ctx context.Context, number string) (*response.BlockResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "BlockchainUseCases.GetBlock")
	defer span.End()

	blockNum := new(big.Int)
	blockNum.SetString(number, 10)

	block, err := m.blockchainRepo.GetBlock(ctx, blockNum)
	if err != nil {
		return nil, err
	}

	txHashes := make([]string, len(block.Transactions))
	for i, hash := range block.Transactions {
		txHashes[i] = hash.Hex()
	}

	return &response.BlockResponse{
		Number:       block.Number.String(),
		Hash:         block.Hash.Hex(),
		ParentHash:   block.ParentHash.Hex(),
		Timestamp:    strconv.FormatUint(block.Timestamp, 10),
		Transactions: txHashes,
	}, nil
}

func (m *Module) GetContractValue(ctx context.Context) (*response.ContractValueResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "BlockchainUseCases.GetContractValue")
	defer span.End()

	value, err := m.blockchainRepo.GetContractValue(ctx)
	if err != nil {
		return nil, err
	}

	return &response.ContractValueResponse{
		Value: value.String(),
	}, nil
}

func (m *Module) SetContractValue(ctx context.Context, req *request.ContractValueRequest) (*response.SetContractValueResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "BlockchainUseCases.SetContractValue")
	defer span.End()

	from := common.HexToAddress(req.From)
	value := big.NewInt(int64(req.Value))

	privateKey, err := helper.StringToECDSA("0f4b9e7651d4df651f30036acdc23ab4ec6108e94f26103eef34aa8e211852c2")
	if err != nil {
		return nil, err
	}

	hash, err := m.blockchainRepo.SetContractValue(ctx, privateKey, from, value)
	if err != nil {
		return nil, err
	}

	return &response.SetContractValueResponse{
		TransactionHash: hash.Hex(),
	}, nil
}
