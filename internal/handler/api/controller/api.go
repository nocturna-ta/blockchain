package controller

import (
	"github.com/gofiber/swagger"
	_ "github.com/nocturna-ta/blockchain/docs"
	"github.com/nocturna-ta/blockchain/internal/usecases"
	"github.com/nocturna-ta/golib/router"
	"time"
)

type API struct {
	prefix         string
	port           uint
	readTimeout    time.Duration
	writeTimeout   time.Duration
	requestTimeout time.Duration
	enableSwagger  bool
	blockchainUc   usecases.BlockchainUseCases
}

type Options struct {
	Prefix         string
	Port           uint
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	RequestTimeout time.Duration
	EnableSwagger  bool
	BlockchainUc   usecases.BlockchainUseCases
}

func New(opts *Options) *API {
	return &API{
		prefix:         opts.Prefix,
		port:           opts.Port,
		readTimeout:    opts.ReadTimeout,
		writeTimeout:   opts.WriteTimeout,
		requestTimeout: opts.RequestTimeout,
		enableSwagger:  opts.EnableSwagger,
		blockchainUc:   opts.BlockchainUc,
	}
}

func (api *API) RegisterRoute() *router.FastRouter {
	myRouter := router.New(&router.Options{
		Prefix:         api.prefix,
		Port:           api.port,
		ReadTimeout:    api.readTimeout,
		WriteTimeout:   api.writeTimeout,
		RequestTimeout: api.requestTimeout,
	})

	if api.enableSwagger {
		myRouter.CustomHandler("GET", "/docs/*", swagger.HandlerDefault, router.MustAuthorized(false))
	}

	myRouter.GET("/health", api.Ping, router.MustAuthorized(false))
	myRouter.Group("/v1", func(v1 *router.FastRouter) {
		v1.GET("/balance/:address", api.GetBalance, router.MustAuthorized(false))
		v1.GET("/transaction/:hash", api.GetTransaction, router.MustAuthorized(false))
		v1.GET("/block/:number", api.GetBlock, router.MustAuthorized(false))
		v1.GET("/contract/value", api.GetContractValue, router.MustAuthorized(false))
		v1.POST("/contract/value", api.SetContractValue, router.MustAuthorized(false))
	})

	return myRouter
}
