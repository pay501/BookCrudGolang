package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string  `json: "name"`
	Author      string  `json: "author"`
	Description string  `json: "description"`
	Price       float64 `json: "price"`
}

var test = []Book{
	{Name: "test1", Author: "test1@mail.com", Description: "test1", Price: 90},
	{Name: "test2", Author: "test2@mail.com", Description: "test2", Price: 90},
	{Name: "test3", Author: "test3@mail.com", Description: "test3", Price: 90},
}

func CreateBook(db *gorm.DB, book *Book) {

	result := db.Create(book)

	if result.Error != nil {
		log.Fatalf("Error creating book: %v", result.Error)
	}

	fmt.Println("Create book successful")
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
