package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

var books = []Book{
	{Name: "test1", Author: "test1@mail.com", Description: "test1", Price: 90},
	{Name: "test2", Author: "test2@mail.com", Description: "test2", Price: 90},
	{Name: "test3", Author: "test3@mail.com", Description: "test3", Price: 90},
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
	db.AutoMigrate(&Book{})

	app := fiber.New()

	//todo Create Book
	/* newBook := Book{Name: "Stary Night", Author: "Van Goh", Description: "Book1", Price: 99.99}
	CreateBook(db, &newBook)
	*/
	//todo Get Book
	book := GetBook(db, 1)

	//todo Update Book
	book.Name = "Learn Golang With PMike"
	book.Author = "PMike"
	UpdateBook(db, book)

	//todo Delete Book
	DeleteBook(db, 1)

	app.Listen(":8080")
}
