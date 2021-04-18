package book

import (
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/category"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string
	Author      string
	Year        int
	IdCategory  uint
	Category    category.Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:IdCategory"`
	Stock       int
	Status      string
}
