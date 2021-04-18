package book

import (
	"time"

	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/category"
	"gorm.io/gorm"
)

type Repository interface {
	InsertBook(book Book) (Book, error)
	UpdateBook(book Book) (Book, error)
	GetAll() []Book
	GetByCategory(idCategory uint) []Book
	GetByYear() []Book
	GetByTitle(title string) *Book
	DeleteBook(book Book) error
	StockUpdate(qty int, book Book) (Book, error)
	BookCategory(name string) *uint
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertBook(book Book) (Book, error) {
	err := r.db.Create(&book)
	if err != nil {
		return book, err.Error
	}
	return book, nil
}

func (r *repository) UpdateBook(book Book) (Book, error) {
	err := r.db.Find(&book).Error
	if err != nil {
		return book, err
	}
	r.db.Save(&book)
	return book, nil
}

func (r *repository) GetAll() []Book {
	var book []Book
	err := r.db.Find(&book)
	if err == nil {
		return nil
	}
	return book
}

func (r *repository) GetByCategory(idCategory uint) []Book {
	var book []Book
	err := r.db.Where("id_category = ?", idCategory).Find(&book)
	if err == nil {
		return nil
	}
	return book
}

func (r *repository) GetByYear() []Book {
	var book []Book
	now := time.Now()
	y := now.Year()
	y2 := y - 2

	err := r.db.Where("year BETWEEN ? AND ?", y2, y).Order("year desc").Find(&book).Error
	if err != nil {
		return nil
	}
	return book
}

func (r *repository) GetByTitle(title string) *Book {
	var book Book
	err := r.db.First(&book, "title = ?", title).Error
	if err == nil {
		return &book
	}
	return nil
}

func (r *repository) DeleteBook(book Book) error {
	err := r.db.Delete(&book)
	if err != nil {
		return err.Error
	}
	return nil
}

func (r *repository) StockUpdate(qty int, book Book) (Book, error) {
	err := r.db.First(&book).Error
	if err != nil {
		return book, err
	}
	book.Stock = book.Stock + qty
	r.db.Save(&book)
	return book, nil
}

func (r *repository) BookCategory(name string) *uint {
	var cat category.Category
	err := r.db.Model(&cat).First(&cat, "name = ?", name).Error
	if err != nil {
		return nil
	}
	return &cat.ID
}
