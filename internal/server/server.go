package server

import (
	"context"
	"github.com/alganbr/kedai-api-gateway/internal/routes"
	"github.com/alganbr/kedai-utils/configs"
	"go.uber.org/fx"
)

func StartApplication(
	lifecycle fx.Lifecycle,
	cfg configs.Config,
	router routes.Router,
	routes routes.Routes,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			routes.Setup()
			router.Run(cfg)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
