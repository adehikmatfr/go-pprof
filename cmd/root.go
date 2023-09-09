package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-pprof",
		Short: "Go Pprof - Service",
	}
)

func Execute() {
	rootCmd.AddCommand(ServeHTTPCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Error: \n", err.Error())
		os.Exit(-1)
	}
}
