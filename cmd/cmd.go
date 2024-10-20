package cmd

import (
	"context"

	"github.com/mamtaharris/risky-plumber/pkg/server"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "risky-plumber",
	Short: "Risky Plumber Service",
}

func Execute(ctx context.Context) error {
	cmd.AddCommand(startServer(ctx))
	if err := cmd.Execute(); err != nil {
		return err
	}

	return nil
}

func startServer(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Start(ctx)
		},
	}
}
