package book

import (
	"time"

	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/category"
)

type ResponseBook struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Author      string            `json:"author"`
	Year        int               `json:"year"`
	IdCategory  uint              `json:"id_category"`
	Category    category.Category `json:"category"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

func BookResponseFormatter(book Book) ResponseBook {
	formatter := ResponseBook{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		Year:        book.Year,
		IdCategory:  book.IdCategory,
		Category:    book.Category,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}
	return formatter
}
