package server

import (
	"context"
	"fmt"

	"github.com/mamtaharris/risky-plumber/config"
	"github.com/mamtaharris/risky-plumber/internal/router"
)

func Start(ctx context.Context) error {
	router, err := router.SetRouter(ctx)
	if err != nil {
		return err
	}
	err = router.Run(":" + fmt.Sprintf("%d", config.App.Port))
	if err != nil {
		return err
	}
	return nil
}
