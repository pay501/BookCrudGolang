package handler

import (
	service "spy/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	newUser := service.User{}
	if err := c.BodyParser(&newUser); err != nil {
		return c.JSON(fiber.Map{
			"message": "wrong with body parser",
			"status":  fiber.StatusBadRequest,
		})
	}

	if newUser.Email == "" || newUser.Password == "" {
		return c.JSON(fiber.Map{
			"message": "Not null",
			"status":  fiber.StatusBadRequest,
		})
	}

	err := service.CreateUser(&newUser)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
			"status":  fiber.StatusForbidden,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Register successful",
		"status":  fiber.StatusCreated,
	})
}
func Login(c *fiber.Ctx) error {
	data := new(service.User)

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
			"status":  fiber.StatusBadRequest,
		})
	}

	if data.Email == "" || data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "data can't be null.",
			"status":  fiber.StatusBadRequest,
		})
	}

	result, err := service.LoginUser(data)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Login failed",
			"status":  fiber.StatusUnauthorized,
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    result,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	}

	c.Cookie((&cookie))
	return c.JSON(fiber.Map{
		"message": "Login successful",
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Microsecond),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
