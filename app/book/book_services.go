package book

import (
	"errors"
)

type Services interface {
	CreateBook(req RequestBook) (Book, error)
	UpdateBook(req RequestBook) (Book, error)
	DeleteBook(req RequestBook) error
	IsBookExist(req RequestBook) error
	NewestBooks() []Book
	BooksByCategory(idCategory uint) []Book
	AllBooks() []Book
	UpdateStock(qty int, req RequestBook) error
	LoanBook(req RequestBook) error
	Return(req RequestBook) error
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) CreateBook(req RequestBook) (Book, error) {
	book := Book{
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
		Year:        req.Year,
	}
	if cat := s.repository.BookCategory(req.Category.Name); cat != nil {
		book.IdCategory = *cat
	} else {
		book.Category = req.Category
	}
	newBook, err := s.repository.InsertBook(book)
	if err != nil {
		return newBook, err
	}
	return newBook, nil
}

func (s *services) UpdateBook(req RequestBook) (Book, error) {
	book := Book{}
	book.Title = req.Title
	book.Description = req.Description
	book.Author = req.Author
	book.Year = req.Year
	book.IdCategory = req.IdCategory
	book.Category = req.Category

	editBook, err := s.repository.UpdateBook(book)
	if err != nil {
		return book, err
	}

	return editBook, nil
}

func (s *services) DeleteBook(req RequestBook) error {
	book := Book{}
	book.Title = req.Title
	book.Description = req.Description
	book.Author = req.Author
	book.Year = req.Year
	book.IdCategory = req.IdCategory
	err := s.repository.DeleteBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (s *services) IsBookExist(req RequestBook) error {
	title := req.Title
	if book := s.repository.GetByTitle(title); book != nil {
		return errors.New("Book already exist")
	}
	return nil
}

func (s *services) NewestBooks() []Book {
	get := s.repository.GetByYear()
	if get != nil {
		return get
	}
	return nil
}

func (s *services) BooksByCategory(idCategory uint) []Book {
	get := s.repository.GetByCategory(idCategory)
	if get != nil {
		return get
	}
	return nil
}

func (s *services) AllBooks() []Book {
	get := s.repository.GetAll()
	if get != nil {
		return get
	}
	return nil
}

func (s *services) UpdateStock(qty int, req RequestBook) error {
	book := Book{}
	book.Title = req.Title
	_, err := s.repository.StockUpdate(qty, book)
	if err != nil {
		return err
	}
	return nil
}

func (s *services) LoanBook(req RequestBook) error {
	book := Book{}
	book.Title = req.Title
	_, err := s.repository.StockUpdate(-1, book)
	if err != nil {
		return err
	}
	return nil
}

func (s *services) Return(req RequestBook) error {
	book := Book{}
	book.Title = req.Title
	_, err := s.repository.StockUpdate(1, book)
	if err != nil {
		return err
	}
	return nil
}
