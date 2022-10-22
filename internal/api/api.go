package api

import "github.com/eduardo-sl/go-archetype-webapp/internal/api/router"

type Api struct {
	Router router.Router
}

func NewAPI() *Api {

	// setup routes
	appRouter := router.NewRouter()

	return &Api{
		Router: *appRouter,
	}
}

func (a *Api) Run() {
	router := a.Router
	router.Echo.Logger.Fatal(router.Echo.Start(":1323"))
}
