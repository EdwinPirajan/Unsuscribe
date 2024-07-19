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
	type request struct {
		Email string `json:"email"`
	}

	var req request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	err := h.Service.Unsubscribe(req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not unsubscribe")
	}

	return c.JSON(http.StatusOK, "Unsubscribed successfully")
}
