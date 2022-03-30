package server

import (
	"github.com/alganbr/kedai-api-gateway/configs"
	"github.com/alganbr/kedai-api-gateway/internal/controllers"
	"github.com/alganbr/kedai-api-gateway/internal/routes"
	authsvc "github.com/alganbr/kedai-authsvc-client/client"
	usersvc "github.com/alganbr/kedai-usersvc-client/client"
	"go.uber.org/fx"
)

var controller = fx.Options(
	fx.Provide(controllers.NewHomeController),
	fx.Provide(controllers.NewUsersController),
)

var router = fx.Options(
	fx.Provide(routes.NewRouter),
	fx.Provide(routes.NewRoutes),
	fx.Provide(routes.NewSwaggerRoutes),
	fx.Provide(routes.NewHomeRoutes),
	fx.Provide(routes.NewUserRoutes),
)

var server = fx.Options(
	fx.Provide(configs.NewConfig),
)

var clients = fx.Options(
	fx.Provide(authsvc.AuthSvcClient),
	fx.Provide(usersvc.UserSvcClient),
)

var Module = fx.Options(
	server,
	router,
	controller,
	fx.Invoke(StartApplication),
)
