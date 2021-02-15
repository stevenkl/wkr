package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/stevenkl/wkr/pkg/config"
	"github.com/stevenkl/wkr/pkg/webserver"
)

var appConfig struct {
	Host          string `env:"WKR_HOST" default:"localhost"`
	Port          int    `env:"WKR_PORT" default:"8000"`
	Data          string `env:"WKR_DATA" default:"./wkr_data"`
	AdminUser     string `env:"WKR_ADMIN_USER:" default:"admin"`
	AdminPassword string `env:"WKR_ADMIN_PASSWORD" default:"admin"`
}

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)

	app := fiber.New()
	app.Use(logger.New())
	webserver.RegisterRoutes(app)

	serverAddress := fmt.Sprintf("%s:%d", config.Host, config.Port)
	fmt.Println(serverAddress)

	app.Listen(serverAddress)
}
