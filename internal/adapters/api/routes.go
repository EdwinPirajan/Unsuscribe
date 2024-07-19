package api

import (
	"github.com/EdwinPirajan/Unsuscribe.git/internal/core/ports"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, service ports.UnsubscribeService) {
	handler := NewHandler(service)
	e.POST("/unsubscribe", handler.Unsubscribe)
}
