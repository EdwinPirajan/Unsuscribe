package main

import (
	"log"

	"github.com/EdwinPirajan/Unsuscribe.git/internal/adapters/api"
	"github.com/EdwinPirajan/Unsuscribe.git/internal/adapters/repository"
	"github.com/EdwinPirajan/Unsuscribe.git/internal/config"
	"github.com/EdwinPirajan/Unsuscribe.git/internal/core/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Inicializar la configuración
	cfg := config.LoadConfig()

	// Inicializar la base de datos
	//repository.InitDB()

	// Inicializar el servidor Echo
	e := echo.New()

	// Configurar el middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Permitir todos los orígenes, puedes ajustar esto según tus necesidades
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Inicializar el repositorio y el servicio
	unsubscribeRepository := repository.NewDBRepository()
	unsubscribeService := services.NewUnsubscribeService(unsubscribeRepository)

	// Configurar rutas
	api.RegisterRoutes(e, unsubscribeService)

	// Iniciar el servidor
	log.Fatal(e.Start(cfg.Server.Address))
}
