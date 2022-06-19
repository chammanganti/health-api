package main

import (
	"log"

	"github.com/chammanganti/health-api/internal/config"
	"github.com/chammanganti/health-api/internal/router"

	"github.com/gofiber/fiber/v2"
)

type App struct{}

func (app *App) Run() error {
	config := config.GetConfig()

	f := fiber.New()

	router.SetUpRoutes(f)

	if err := f.Listen(config.ADDR); err != nil {
		log.Fatal("Failed setting up the server")
		return err
	}

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
