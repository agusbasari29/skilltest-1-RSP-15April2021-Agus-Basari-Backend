package loan

import (
	"time"

	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/book"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/user"
)

type ResponseLoan struct {
	IdUser       int       `json:"id_user"`
	User         user.User `json:"user"`
	IdBook       int       `json:"id_book"`
	Book         book.Book `json:"book"`
	BorrowedDate time.Time `json:"borrowed_date"`
	DueDate      time.Time `json:"due_date"`
}

func LoanResponseFormatter(loan Loan) ResponseLoan {
	formatter := ResponseLoan{
		IdUser:       loan.IdUser,
		User:         loan.User,
		IdBook:       loan.IdBook,
		Book:         loan.Book,
		BorrowedDate: loan.BorrowedDate,
		DueDate:      loan.DueDate,
	}
	return formatter
}
