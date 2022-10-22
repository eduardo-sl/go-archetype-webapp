package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthcheckHandler struct {
}

func NewHealthCheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) Handle(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, response{Status: pass})
}