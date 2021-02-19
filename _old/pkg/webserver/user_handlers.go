package webserver

import (
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/stevenkl/wkr/pkg/models"
	"github.com/stevenkl/wkr/pkg/service"
)

func userLoginHandler(c *fiber.Ctx) error {
	user := new(models.UserModel)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var userFromDB models.UserModel
	for _, u := range serverConfig.Users {
		if u.Name == user.Name {
			userFromDB = u
		}
	}
	if userFromDB.Name == "" {
		// User not found
		return c.SendStatus(fiber.StatusForbidden)
	}

	if service.CheckPasswordHash(user.PasswordHash, userFromDB.PasswordHash) {
		// hash matches
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = userFromDB.Name
		claims["group"] = userFromDB.Group
		claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"status": "success",
			"data":   t,
		})
	}

	// password hash doesnt match
	return c.SendStatus(fiber.StatusUnauthorized)
	
}
