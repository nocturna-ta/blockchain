package cmd

import (
	"github.com/nocturna-ta/golib/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Blockchain Service",
		Short: "Blockchain Service",
	}
)

func Execute() {
	log.SetFormatter("json")
	rootCmd.AddCommand(nil)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Error: ", err.Error())
		os.Exit(-1)
	}
}