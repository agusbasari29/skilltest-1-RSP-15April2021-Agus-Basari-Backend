package payment

import (
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/loan"
	"gorm.io/gorm"
)

type FinePayment struct {
	gorm.Model
	Receipt string
	Amount  float32
	IdLoan  int
	Loan    loan.Loan `gorm:"foreignKey:IdLoan"`
}
