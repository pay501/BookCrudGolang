package handler

import (
	"spy/errs"
	"spy/repository"
	service "spy/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h *userHandler) SignUp(c *fiber.Ctx) error {
	newUser := repository.User{}
	if err := c.BodyParser(&newUser); err != nil {
		return c.JSON(fiber.Map{
			"result": errs.NewBadRequestError("bad request"),
		})
	}

	err := h.userService.SignUp(&newUser)
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
func (h *userHandler) Login(c *fiber.Ctx) error {
	data := service.LoginReq{}

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
			"status":  fiber.StatusBadRequest,
		})
	}

	result, err := h.userService.Login(&data)
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
		"token":   result,
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
