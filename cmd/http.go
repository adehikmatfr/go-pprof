package cmd

import (
	"github.com/adehikmatfr/go-pprof/handler"
	"github.com/spf13/cobra"
)

var (
	serveHTTPCmd = &cobra.Command{
		Use:   "serve-http",
		Short: "Centralize go-pprof Service",
		Long:  "Centralize go-pprof Service",
		RunE:  runHTTP,
	}
)

func ServeHTTPCmd() *cobra.Command {
	return serveHTTPCmd
}

func runHTTP(cmd *cobra.Command, args []string) error {
	bootstrapHTTP()
	return nil
}

func bootstrapHTTP() {
	handler.RunServeHttp()
}
