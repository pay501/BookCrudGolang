package main

import (
	"fmt"
	"log"
	"os"
	"spy/handler"
	service "spy/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

var Db *gorm.DB

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	// db.Migrator().DropColumn(&Book{}, "test")
	Db.AutoMigrate(&service.Book{}, &service.User{})

	service.Db = *Db

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
	app.Post("/register", handler.SignUp)
	app.Post("/login", handler.Login)
	app.Get("/logout", handler.Logout)
	app.Get("/books", handler.GetBooks)

	//todo Use middleware here
	app.Use(AuthRequestMiddleware)
	app.Get("/checkSession", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Session true",
		})
	})

	// controller here
	app.Get("/book/:id", handler.GetBookById)
	app.Post("/addBook", handler.CreateBook)
	app.Put("/updateBook/:id", handler.UpdateBook)
	app.Delete("/deleteBook/:id", handler.DeleteBook)

	app.Listen(":8080")
}

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold: time.Second, // Slow SQL threshold
		LogLevel:      logger.Info, // Log level
		Colorful:      true,        // Disable color
	},
)
