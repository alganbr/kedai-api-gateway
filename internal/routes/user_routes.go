package routes

import "github.com/alganbr/kedai-api-gateway/internal/controllers"

type UserRoutes struct {
	router          Router
	usersController controllers.IUsersController
}

func NewUserRoutes(router Router, usersController controllers.IUsersController) UserRoutes {
	return UserRoutes{
		router:          router,
		usersController: usersController,
	}
}

func (r *UserRoutes) Setup() {
	userGroup := r.router.Path.Group("/users")
	userGroup.Use(r.router.AuthRequired)
	userGroup.GET("/:id", r.usersController.Get)
	userGroup.POST("", r.usersController.Create)
	userGroup.PUT("/:id", r.usersController.Update)
	userGroup.PATCH("/:id", r.usersController.Patch)
}
