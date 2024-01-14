package main

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
}

func CreateBook(db *gorm.DB, book *Book) error {

	result := db.Create(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetBooks(db *gorm.DB) []Book {
	var books []Book
	result := db.Find(&books)

	if result.Error != nil {
		log.Fatalf("Error getting books: %v", result.Error)
	}

	return books
}

func GetBook(db *gorm.DB, id int) (*Book, error) {
	var book Book

	result := db.First(&book, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func UpdateBook(db *gorm.DB, book *Book, id int) error {
	result := db.Where("id = ?", id).Save(book)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Updated book successfully")

	return nil
}

func DeleteBook(db *gorm.DB, id int) error {
	var book Book

	result := db.Delete(&book, id)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Delete book successful")

	return nil
}
