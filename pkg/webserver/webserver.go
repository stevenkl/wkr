package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/stevenkl/wkr/pkg/config"
)

var serverConfig *config.AppConfig

func New(config *config.AppConfig) *fiber.App {
	app := fiber.New()
	serverConfig = config

	app.Use(logger.New())
	app.Get("/ping", statusPingHandler)
	app.Post("/login", userLoginHandler)

	jobsApp := fiber.New()
	jobsApp.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	jobsApp.Get("/", jobsIndexHandler)
	jobsApp.Get("/:job_id", jobsDetailsHandler)
	jobsApp.Get("/:job_id/:run_id", jobExecutionsHandler)

	app.Mount("/jobs", jobsApp)

	return app
}
