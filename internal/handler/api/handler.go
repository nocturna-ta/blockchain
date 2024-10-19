package api

import (
	"github.com/nocturna-ta/blockchain/config"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/router"
)

type Options struct {
	Cfg config.MainConfig
}

type Handler struct {
	opts        *Options
	listenErrCh chan error
	myRouter    *router.FastRouter
}

func New(opts *Options) *Handler {
	handler := &Handler{
		opts: opts,
	}
	handler.myRouter = contr
}
func (h *Handler) Run() {
	log.Infof("API Listening on %d", h.opts.Cfg.Server.Port)
	h.listenErrCh <- h.myRouter.StartServe()
}

func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}
