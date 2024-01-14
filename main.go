package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Second, // Slow SQL threshold
		LogLevel:      logger.Info, // Log level
		Colorful:      true,        // Disable color
	},
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	// db.Migrator().DropColumn(&Book{}, "test")
	db.AutoMigrate(&Book{})

	app := fiber.New()

	//todo Get Books
	app.Get("/books", corsMiddleware, func(c *fiber.Ctx) error {
		return c.JSON(GetBooks(db))
	})

	app.Get("/book/:id", corsMiddleware, func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.JSON(fiber.Map{
				"message": "Not Found",
				"status":  fiber.StatusNotFound,
			})
		}
		book, err := GetBook(db, id)

		if err != nil {
			c.JSON(fiber.Map{
				"error": err,
			})
		}
		return c.JSON(fiber.Map{
			"result": book,
		})
	})

	//todo Create Book
	app.Post("/addBook", corsMiddleware, func(c *fiber.Ctx) error {
		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return c.JSON(fiber.Map{
				"message": "There is something wrong",
				"status":  fiber.StatusBadRequest,
			})
		}

		CreateBook(db, book)

		return c.JSON(fiber.Map{
			"book":   book,
			"status": fiber.StatusCreated,
		})
	})

	//todo Update Book
	app.Put("/updateBook/:id", corsMiddleware, func(c *fiber.Ctx) error {
		book := new(Book)
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

		UpdateBook(db, book, id)
		return c.JSON(fiber.Map{
			"message": book,
			"status":  fiber.StatusOK,
		})
	})

	//todo Delete
	app.Delete("/deleteBook/:id", corsMiddleware, func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.JSON(fiber.Map{
				"message": err,
			})
		}

		DeleteBook(db, id)
		return c.JSON(fiber.Map{
			"message": "delete successful",
		})
	})

	app.Listen(":8080")
}

func corsMiddleware(c *fiber.Ctx) error {
	// Enable CORS for all routes
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")

	// Continue to next middleware or route handler
	return c.Next()
}
