package handler

import (
	"fmt"
	"net/http"
	service "spy/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	books := service.GetBooks()
	return c.JSON(fiber.Map{
		"code":   http.StatusOK,
		"result": books,
	})
}

func GetBookById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Not Found",
			"status":  fiber.StatusNotFound,
		})
	}

	book, err := service.GetBook(id)
	if err != nil {
		c.JSON(fiber.Map{
			"error": err,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Get book successful",
		"result":  book,
	})
}

func CreateBook(c *fiber.Ctx) error {
	book := new(service.Book)
	if err := c.BodyParser(book); err != nil {
		return c.JSON(fiber.Map{
			"message": "There is something wrong",
			"status":  fiber.StatusBadRequest,
		})
	}

	if book.Name == "" || book.Author == "" || book.Description == "" || book.Price == 0 {
		return c.JSON(fiber.Map{
			"message": "Data is null.",
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(book.Image)
	err := service.CreateBook(book)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
		})
	}

	return c.JSON(fiber.Map{
		"message": "created successful",
		"status":  fiber.StatusCreated,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	book := new(service.Book)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
			"status":  fiber.StatusNotFound,
		})
	}
	if err = c.BodyParser(book); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
			"status":  fiber.StatusBadRequest,
		})
	}

	err = service.UpdateBook(book, id)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
			"status":  fiber.StatusBadRequest,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Update successful",
		"status":  fiber.StatusOK,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	service.DeleteBook(id)
	return c.JSON(fiber.Map{
		"message": "delete successful",
	})
}
