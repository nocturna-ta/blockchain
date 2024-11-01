package usecases

import (
	"context"
	"github.com/nocturna-ta/blockchain/internal/usecases/request"
	"github.com/nocturna-ta/blockchain/internal/usecases/response"
)

type BlockchainUseCases interface {
	GetBalance(ctx context.Context, address string) (*response.BalanceResponse, error)
	GetTransaction(ctx context.Context, hash string) (*response.TransactionResponse, error)
	GetBlock(ctx context.Context, number string) (*response.BlockResponse, error)
	GetContractValue(ctx context.Context) (*response.ContractValueResponse, error)
	SetContractValue(ctx context.Context, req *request.ContractValueRequest) (*response.SetContractValueResponse, error)
}
