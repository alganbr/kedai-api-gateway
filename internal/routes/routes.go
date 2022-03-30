package routes

type Routes struct {
	routes []Route
}

type Route interface {
	Setup()
}

func NewRoutes(
	swaggerRoutes SwaggerRoutes,
	homeRoutes HomeRoutes,
	authRoutes AuthRoutes,
	userRoutes UserRoutes,
) Routes {
	return Routes{
		routes: []Route{
			&swaggerRoutes,
			&homeRoutes,
			&authRoutes,
			&userRoutes,
		},
	}
}

func (r *Routes) Setup() {
	for _, route := range r.routes {
		route.Setup()
	}
}
