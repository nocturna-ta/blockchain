package controller

import (
	"context"
	"encoding/json"
	"github.com/nocturna-ta/blockchain/internal/usecases/request"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/golib/tracing"
)

// GetBalance godoc
// @Summary 	Get Balance
// @Description	Get Balance
// @Tags		blockchain
// @Accept		json
// @Param		X-Channel-Id			header		string	false 	"channel where request comes from"	default(web)
// @Param		address						path		string 	true	"account address"
// @Produce		json
// @Success		200	{object}	jsonResponse{data=response.BalanceResponse}
// @Router		/v1/balance/{address}	[get]
func (api *API) GetBalance(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetBalance")
	defer span.End()

	res, err := api.blockchainUc.GetBalance(ctx, req.Params("address"))
	if err != nil {
		return nil, err
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetTransaction godoc
// @Summary 	Get Transaction
// @Description	Get Transaction
// @Tags		blockchain
// @Accept		json
// @Param		X-Channel-Id			header		string	false 	"channel where request comes from"	default(web)
// @Param		hash						path		string 	true	"transaction hash"
// @Produce		json
// @Success		200	{object}	jsonResponse{data=response.TransactionResponse}
// @Router		/v1/transaction/{hash}	[get]
func (api *API) GetTransaction(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetTransaction")
	defer span.End()

	res, err := api.blockchainUc.GetTransaction(ctx, req.Params("hash"))
	if err != nil {
		return nil, err
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetBlock godoc
// @Summary 	Get Block
// @Description	Get Block
// @Tags		blockchain
// @Accept		json
// @Param		X-Channel-Id			header		string	false 	"channel where request comes from"	default(web)
// @Param		number						path		string 	true	"block number"
// @Produce		json
// @Success		200	{object}	jsonResponse{data=response.BlockResponse}
// @Router		/v1/block/{number}	[get]
func (api *API) GetBlock(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetBlock")
	defer span.End()

	res, err := api.blockchainUc.GetBlock(ctx, req.Params("number"))
	if err != nil {
		return nil, err
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetContractValue godoc
// @Summary 	Get Contract Value
// @Description	Get Contract Value
// @Tags		blockchain
// @Accept		json
// @Param		X-Channel-Id			header		string	false 	"channel where request comes from"	default(web)
// @Produce		json
// @Success		200	{object}	jsonResponse{data=response.ContractValueResponse}
// @Router		/v1/contract/value	[get]
func (api *API) GetContractValue(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetContractValue")
	defer span.End()

	res, err := api.blockchainUc.GetContractValue(ctx)
	if err != nil {
		return nil, err
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// SetContractValue godoc
// @Summary 	Create Contract Value
// @Description Create Contract Value
// @Tags		blockchain
// @Accept		json
// @Param		X-Channel-Id	header		string	false 	"channel where request comes from"	default(web)
// @Param		contracts		body 		request.ContractValueRequest	true	"Contract Value"
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.SetContractValueResponse}
// @Router		/v1/contract/value [post]
func (api *API) SetContractValue(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.SetContractValue")
	defer span.End()

	var contractValueReq request.ContractValueRequest
	err := json.Unmarshal(req.RawBody(), &contractValueReq)
	if err != nil {
		return nil, err
	}

	res, err := api.blockchainUc.SetContractValue(ctx, &contractValueReq)
	if err != nil {
		return nil, err
	}

	return rest.NewJSONResponse().SetData(res), nil
}
