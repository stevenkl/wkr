package webserver

import "github.com/gofiber/fiber/v2"

func statusPingHandler(c *fiber.Ctx) error {
	return c.SendString("Pong")
}
