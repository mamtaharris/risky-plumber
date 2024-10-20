package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/mamtaharris/risky-plumber/cmd"
	"github.com/mamtaharris/risky-plumber/config"
	"github.com/mamtaharris/risky-plumber/pkg/logger"
)

func main() {
	config.InitConfig()
	logger.InitLogger()
	ctx := context.Background()
	err := cmd.Execute(ctx)
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
	go func() {
		_, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
		defer stop()
	}()
}
