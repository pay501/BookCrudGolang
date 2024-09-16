package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthRequestMiddleware(c *fiber.Ctx) error {
	jwtSecretKey := "testSecret"
	cookie := c.Cookies("jwt")

	if cookie == "" {
		c.Redirect("/redirect")
		return nil
	}

	_, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		c.Redirect("/redirect")
	}

	return c.Next()
}
