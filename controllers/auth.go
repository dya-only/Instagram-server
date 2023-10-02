package controllers

import (
	_user "go-template/ent/user"
	"go-template/models"
	"go-template/utils"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginByPassword(c *fiber.Ctx) error {
	var body models.User

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	id, _ := utils.DbConn.User.Query().Where(_user.EmailEQ(body.Email)).OnlyID(ctx)

	if id == 0 {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Not found user",
		})
	}

	user, _ := utils.DbConn.User.Get(ctx, id)

	if body.Password != user.Password {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Failed to login by password",
		})
	}

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := _token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}

func LoginByToken(c *fiber.Ctx) error {
	var body models.Verify

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(body.Token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Invalid token",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"claims":  token.Claims,
	})
}
