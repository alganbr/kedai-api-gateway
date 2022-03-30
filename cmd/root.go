package cmd

import (
	"github.com/alganbr/kedai-api-gateway/internal/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var RootCmd = &cobra.Command{
	Use:   "api-gateway",
	Short: "API Gateway",
	Long:  "API Gateway",
	RunE: func(cmd *cobra.Command, args []string) error {
		fx.New(server.Module).Run()
		return nil
	},
}
