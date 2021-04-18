package book

import "github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/category"

type RequestBook struct {
	Title       string            `json:"title" validate:"required"`
	Description string            `json:"description" validate:"required"`
	Author      string            `json:"author" validate:"required"`
	Year        int               `json:"year" validate:"required"`
	IdCategory  uint              `json:"id_category"`
	Category    category.Category `json:"category"`
}
