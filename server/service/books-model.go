package model

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

func CreateBook(book *Book) error {

	result := Db.Create(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetBooks() []Book {
	var books []Book
	result := Db.Find(&books)

	if result.Error != nil {
		log.Fatalf("Error getting books: %v", result.Error)
	}

	return books
}

func GetBook(id int) (*Book, error) {
	var book Book

	result := Db.First(&book, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func UpdateBook(book *Book, id int) error {

	result := Db.Model(&book).Where("id = ?", id).Updates(book)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Updated book successfully")

	return nil
}

func DeleteBook(id int) error {
	var book Book

	result := Db.Delete(&book, id)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Delete book successful")

	return nil
}
