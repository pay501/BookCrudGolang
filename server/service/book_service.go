package model

import "spy/repository"

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{bookRepo: bookRepo}
}

func (s *bookService) AddBook(book *repository.Book) error {
	err := s.bookRepo.CreateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (s *bookService) GetSingleBook(id int) (*repository.Book, error) {
	result, err := s.bookRepo.GetBook(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *bookService) GetAllBook() ([]repository.Book, error) {
	books, err := s.bookRepo.GetBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *bookService) UpdateBookService(book *repository.Book, id int) error {
	err := s.bookRepo.UpdateBook(book, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *bookService) DeleteBookService(id int) error {
	err := s.bookRepo.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}
