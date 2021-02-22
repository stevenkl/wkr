package main


import (
	"time"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwt "github.com/form3tech-oss/jwt-go"
	jwtware "github.com/gofiber/jwt/v2"
)

func registerAppHandlers() {

	app.Use(logger.New())

	app.Get("/", indexHandler)

	app.Post("/login", loginHandler)

	secret := config.Server.Secret.String()
	jobs := fiber.New()
	jobs.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	}))

	jobs.Get("/", jobsIndexHandler)
	jobs.Get("/:job_id", jobsDetailsHandler)

	app.Mount("/jobs", jobs)
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
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)

		claims["name"] = configuser.Name
		claims["group"] = configuser.Group
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		secret := config.Server.Secret.String()
		t, err := token.SignedString([]byte(secret))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		
		return c.JSON(fiber.Map{
			"status": "success",
			"data": t,
		})
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}


func jobsIndexHandler(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status": "success",
		"data": config.Jobs.Jobs,
	})
}

func jobsDetailsHandler(c *fiber.Ctx) error {
	search := c.Params("job_id")
	job, err := config.Jobs.GetByID(search)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data": job,
	})
}