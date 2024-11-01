package server

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nocturna-ta/blockchain/config"
	"github.com/nocturna-ta/blockchain/internal/handler/api"
	"github.com/nocturna-ta/golib/log"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var (
	serverHTTPCmd = &cobra.Command{
		Use:   "server-http",
		Short: "Blockchain Service HTTP",
		Long:  "Blockchain Service HTTP",
		RunE:  run,
	}
)

func ServeHttpCmd() *cobra.Command {
	serverHTTPCmd.Flags().StringP("config", "c", "", "Config Path, both relative or absolute. i.e: /usr/local/bin/config/files")
	return serverHTTPCmd
}

func run(cmd *cobra.Command, args []string) error {
	configLocation, _ := cmd.Flags().GetString("config")
	cfg := &config.MainConfig{}
	config.ReadConfig(cfg, configLocation)

	client, err := ethclient.Dial(cfg.Blockchain.GanacheURL)
	if err != nil {
		return err
	}
	appContainer := newContainer(&options{
		Cfg:    cfg,
		Client: client,
	})

	server := api.New(&api.Options{
		Cfg:          appContainer.Cfg,
		BlockchainUc: appContainer.BlockchainUc,
	})

	go server.Run()

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Info("Exiting gracefully...")
	case err := <-server.ListenError():
		log.Error("Error starting web server, exiting gracefully:", err)
	}

	return nil
}
