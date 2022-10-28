package router

import (
	"github.com/eduardo-sl/go-webapp-project/internal/api/handlers/healthcheck"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo *echo.Echo
}

func NewRouter() *Router {
	e := echo.New()

	setupHealthRoutes(e)

	return &Router{
		Echo: e,
	}
}

func setupHealthRoutes(e *echo.Echo) {
	e.GET("/healthcheck", healthcheck.NewHealthCheckHandler().Handle)
}
