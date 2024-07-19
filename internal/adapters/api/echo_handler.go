package api

import (
	"net/http"

	"github.com/EdwinPirajan/Unsuscribe.git/internal/core/ports"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service ports.UnsubscribeService
}

func NewHandler(service ports.UnsubscribeService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) Unsubscribe(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, "Email is required")
	}

	err := h.Service.Unsubscribe(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not unsubscribe")
	}

	return c.Redirect(http.StatusFound, "https://www.bancoagrario.gov.co")
}
