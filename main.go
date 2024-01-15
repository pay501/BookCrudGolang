package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

func corsMiddleware(c *fiber.Ctx) error {
	// Enable CORS for all routes
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type")

	// Continue to next middleware or route handler
	return c.Next()
}

func authRequired(c *fiber.Ctx) error {
	jwtSecretKey := "testSecret"
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claim := token.Claims.(jwt.MapClaims)
	fmt.Println(claim["user_id"])
	return c.Next()
}

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
	db.AutoMigrate(&Book{}, &User{})

	app := fiber.New()
	app.Use("/books", authRequired)

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

		err = CreateBook(db, book)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
			})
		}

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

		err = UpdateBook(db, book, id)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err,
				"status":  fiber.StatusBadRequest,
			})
		}
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

	//todo User API
	app.Post("/register", corsMiddleware, func(c *fiber.Ctx) error {
		newUser := new(User)
		if err := c.BodyParser(&newUser); err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
			})
		}

		if newUser.Email == "" || newUser.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Not null",
				"status":  fiber.StatusBadRequest,
			})
		}

		err = CreateUser(db, newUser)
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
	})

	app.Post("/login", corsMiddleware, func(c *fiber.Ctx) error {
		data := new(User)

		if err := c.BodyParser(data); err != nil {
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

		result, err := LoginUser(db, data)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    result,
			Expires:  time.Now().Add(time.Hour * 72),
			HTTPOnly: true,
		})

		return c.JSON(fiber.Map{
			"message": "Login successful.",
		})
	})

	app.Listen(":8080")
}
