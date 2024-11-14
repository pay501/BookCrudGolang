package handler

import (
	"net/http"
	"spy/repository"
	service "spy/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type bookHandler struct {
	bookServive service.BookService
}

func NewBookHandler(bookService service.BookService) bookHandler {
	return bookHandler{bookServive: bookService}
}

func (h *bookHandler) GetBooks(c *fiber.Ctx) error {
	books, err := h.bookServive.GetAllBook()
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": "cannot get books",
		})
	}
	return c.JSON(fiber.Map{
		"code":   http.StatusOK,
		"result": books,
	})
}

func (h *bookHandler) GetBookById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Not Found",
			"status":  fiber.StatusNotFound,
		})
	}

	book, err := h.bookServive.GetSingleBook(id)
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

func (h *bookHandler) CreateBook(c *fiber.Ctx) error {
	// or book := new(repository.Book)
	book := repository.Book{}
	if err := c.BodyParser(&book); err != nil {
		return c.JSON(fiber.Map{
			"message": "There is something wrong",
			"status":  fiber.StatusBadRequest,
		})
	}

	err := h.bookServive.AddBook(&book)
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

func (h *bookHandler) UpdateBook(c *fiber.Ctx) error {
	book := new(repository.Book)
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

	err = h.bookServive.UpdateBookService(book, id)
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

func (h *bookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	err = h.bookServive.DeleteBookService(id)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}
	return c.JSON(fiber.Map{
		"message": "delete successful",
	})
}
