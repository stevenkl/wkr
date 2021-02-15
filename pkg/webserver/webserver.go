package webserver

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/stevenkl/wkr/pkg/models"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(app *fiber.App) error {
	app.Get("/api/ping", pingHandler)
	app.Post("/api/login", loginHandler)
	app.Get("/api/jobs", jobsHandler)

	return nil
}

func pingHandler(c *fiber.Ctx) error {
	return c.SendString("Pong")
}

func loginHandler(c *fiber.Ctx) error {
	admin := new(models.AdminModel)
	if err := c.BodyParser(admin); err != nil {
		return err
	}
	return c.SendString(fmt.Sprintf("Hello, %s!", admin.Name))
}

func jobsHandler(c *fiber.Ctx) error {
	return c.SendString("Listing all jobs.")
}
