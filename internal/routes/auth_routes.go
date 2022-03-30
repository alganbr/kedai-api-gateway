package routes

import "github.com/alganbr/kedai-api-gateway/internal/controllers"

type AuthRoutes struct {
	router         Router
	authController controllers.IAuthController
}

func NewAuthRoutes(router Router, authController controllers.IAuthController) AuthRoutes {
	return AuthRoutes{
		router:         router,
		authController: authController,
	}
}

func (r *AuthRoutes) Setup() {
	authGroup := r.router.Path.Group("/auth")
	authGroup.GET("/:id", r.authController.Get)
	authGroup.POST("", r.authController.Authenticate)
}
