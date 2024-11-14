package repository

import "gorm.io/gorm"

type bookRepositoryDB struct {
	db *gorm.DB
}

func NewBookRepositoryDB(db *gorm.DB) BookRepository {
	return &bookRepositoryDB{db: db}
}

func (r *bookRepositoryDB) CreateBook(book *Book) error {
	result := r.db.Create(book)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *bookRepositoryDB) GetBook(id int) (*Book, error) {
	var book Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (r *bookRepositoryDB) GetBooks() ([]Book, error) {
	books := []Book{}
	result := r.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (r *bookRepositoryDB) UpdateBook(book *Book, id int) error {
	result := r.db.Model(&book).Where("id = ?", id).Updates(book)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *bookRepositoryDB) DeleteBook(id int) error {
	var book Book
	result := r.db.Delete(&book, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
