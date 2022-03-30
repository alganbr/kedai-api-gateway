package routes

import (
	"github.com/alganbr/kedai-api-gateway/configs"
	"github.com/alganbr/kedai-api-gateway/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type SwaggerRoutes struct {
	router Router
	cfg    *configs.Config
}

func NewSwaggerRoutes(router Router, cfg *configs.Config) SwaggerRoutes {
	return SwaggerRoutes{
		router: router,
		cfg:    cfg,
	}
}

func (r *SwaggerRoutes) Setup() {
	docs.SwaggerInfo.BasePath = "/api"
	r.router.Path.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
