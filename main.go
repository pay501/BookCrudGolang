package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

func authRequestMiddleware(c *fiber.Ctx) error {
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
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/redirect", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "return",
		})
	})

	//! User API

	//? User register
	app.Post("/register", func(c *fiber.Ctx) error {
		newUser := new(User)
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

	//? User Login
	app.Post("/login", func(c *fiber.Ctx) error {
		data := new(User)

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

		result, err := LoginUser(db, data)
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
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
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
	})

	//todo Get Books
	app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Get book successful",
			"result":  GetBooks(db),
		})
	})


	//todo Use middleware here
	app.Use(authRequestMiddleware)

	app.Get("/checkSession", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Session true",
		})
	})

	app.Get("/book/:id", func(c *fiber.Ctx) error {
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
			"message": "Get book successful",
			"result":  book,
		})
	})

	app.Post("/addBook", func(c *fiber.Ctx) error {
		book := new(Book)
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
		err = CreateBook(db, book)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
			})
		}

		return c.JSON(fiber.Map{
			"message": "Add successful",
			"status":  fiber.StatusCreated,
		})
	})

	app.Put("/updateBook/:id", func(c *fiber.Ctx) error {
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
			"message": "Update successful",
			"status":  fiber.StatusOK,
		})
	})

	//todo Delete
	app.Delete("/deleteBook/:id", func(c *fiber.Ctx) error {
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
