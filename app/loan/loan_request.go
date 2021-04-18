package loan

import (
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/book"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/user"
)

type RequestLoan struct {
	ID     uint      `json:"id"`
	IdUser int       `json:"is_user"`
	User   user.User `json:"user"`
	IdBook int       `json:"id_book"`
	Book   book.Book `json:"book"`
}
