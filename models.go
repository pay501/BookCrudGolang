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

func CreateBook(db *gorm.DB, book *Book) {

	result := db.Create(book)

	if result.Error != nil {
		log.Fatalf("Error creating book: %v", result.Error)
	}

	fmt.Println("Create book successful")
}

func GetBooks(db *gorm.DB) []Book {
	var books []Book
	result := db.Find(&books)

	if result.Error != nil {
		log.Fatalf("Error getting books: %v", result.Error)
	}

	return books
}

func GetBook(db *gorm.DB, id uint64) *Book {
	var book Book

	result := db.First(&book, id)

	if result.Error != nil {
		log.Fatalf("Error getting book: %v", result.Error)
	}
	return &book
}

func UpdateBook(db *gorm.DB, book *Book) {
	result := db.Save(book)

	if result.Error != nil {
		log.Fatalf("Error updating book: %v", result.Error)
	}

	fmt.Println("Updated book successfully")
}

func DeleteBook(db *gorm.DB, id uint64) {
	var book Book

	result := db.Delete(&book, id)

	if result.Error != nil {
		log.Fatalf("Error deleting book: %v", result.Error)
	}

	fmt.Println("Delete book successful")
}
