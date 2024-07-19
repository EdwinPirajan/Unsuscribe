package main

import (
	"github.com/EdwinPirajan/Unsuscribe.git/internal/adapters/api"
	"github.com/EdwinPirajan/Unsuscribe.git/internal/adapters/repository"
	"github.com/EdwinPirajan/Unsuscribe.git/internal/config"
	"github.com/EdwinPirajan/Unsuscribe.git/internal/core/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	cfg := config.LoadConfig()

	repository.InitDB()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Permitir todos los orígenes, puedes ajustar esto según tus necesidades
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	unsubscribeRepository := repository.NewDBRepository()
	unsubscribeService := services.NewUnsubscribeService(unsubscribeRepository)

	api.RegisterRoutes(e, unsubscribeService)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(cfg.Server.Address))
}
