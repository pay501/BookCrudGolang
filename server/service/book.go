package model

import (
	"spy/repository"

	"gorm.io/gorm"
)

type BookResponse struct {
	gorm.Model
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

type BookService interface {
	AddBook(book *repository.Book) error
	GetSingleBook(id int) (*repository.Book, error)
	GetAllBook() ([]repository.Book, error)
	UpdateBookService(book *repository.Book, id int) error
	DeleteBookService(id int) error
}
