package webserver

import (
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/stevenkl/wkr/pkg/models"
)

func userLoginHandler(c *fiber.Ctx) error {
	admin := new(models.AdminModel)
	if err := c.BodyParser(admin); err != nil {
		return err
	}

	if admin.Name != serverConfig.AdminUser || admin.Password != serverConfig.AdminPassword {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Administrator"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   t,
	})
}
