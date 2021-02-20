package main


import (
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func registerAppHandlers() {

	app.Use(logger.New())

	app.Get("/", indexHandler)

	app.Post("/login", loginHandler)
}


func indexHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func loginHandler(c *fiber.Ctx) error {
	user := new(UserConfig)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	
	configuser, err := config.Users.GetByName(user.Name)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if validatePasswordHash(user.Password, configuser.Password) {
		// generating JWT Token here and send back
		return c.SendString("Authorized")
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}
