package repository

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

type BookRepository interface {
	CreateBook(book *Book) error
	GetBook(id int) (*Book, error)
	GetBooks() ([]Book, error)
	UpdateBook(book *Book, id int) error
	DeleteBook(id int) error
}
