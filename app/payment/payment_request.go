package payment

import "github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/loan"

type RequestFinePayment struct {
	Receipt string  `json:"receipt"`
	Amount  float32 `json:"amount"`
	IdLoan  int     `json:"id_loan"`
	Loan    loan.Loan
}
