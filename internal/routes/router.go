package routes

import (
	"fmt"
	"github.com/alganbr/kedai-api-gateway/configs"
	_ "github.com/alganbr/kedai-api-gateway/docs"
	"github.com/alganbr/kedai-authsvc-client/client"
	"github.com/alganbr/kedai-utils/datetime"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	Gin     *gin.Engine
	Path    *gin.RouterGroup
	authSvc client.IAuthSvcClient
}

func NewRouter(authSvc client.IAuthSvcClient) Router {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	group := router.Group(fmt.Sprintf("/api"))

	return Router{
		Gin:     router,
		Path:    group,
		authSvc: authSvc,
	}
}

func (r Router) Run(cfg *configs.Config) {
	go func() {
		err := r.Gin.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
		if err != nil {
			panic(err)
		}
	}()
}

func (r Router) AuthRequired(c *gin.Context) {
	accessTokenId := c.Request.Header.Get("Authorization")
	if accessTokenId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Unauthorized",
		})
		return
	}
	accessToken, err := r.authSvc.Auth().Get(accessTokenId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
	if accessToken.Expires < datetime.GetUtcNow().Unix() {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Unauthorized",
		})
		return
	}
	c.Next()
}
