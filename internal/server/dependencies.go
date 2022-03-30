package server

import (
	"github.com/alganbr/kedai-api-gateway/configs"
	"github.com/alganbr/kedai-api-gateway/internal/clients"
	"github.com/alganbr/kedai-api-gateway/internal/controllers"
	"github.com/alganbr/kedai-api-gateway/internal/routes"
	"go.uber.org/fx"
)

var controller = fx.Options(
	fx.Provide(controllers.NewHomeController),
	fx.Provide(controllers.NewAuthController),
	fx.Provide(controllers.NewUsersController),
)

var router = fx.Options(
	fx.Provide(routes.NewRouter),
	fx.Provide(routes.NewRoutes),
	fx.Provide(routes.NewSwaggerRoutes),
	fx.Provide(routes.NewHomeRoutes),
	fx.Provide(routes.NewAuthRoutes),
	fx.Provide(routes.NewUserRoutes),
)

var server = fx.Options(
	fx.Provide(configs.NewConfig),
)

var client = fx.Options(
	fx.Provide(clients.NewAuthSvcClient),
	fx.Provide(clients.NewUserSvcClient),
)

var Module = fx.Options(
	server,
	client,
	router,
	controller,
	fx.Invoke(StartApplication),
)
