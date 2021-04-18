package loan

import (
	"time"

	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/book"
	"github.com/agusbasari29/skilltest-1-RSP-15April2021-Agus-Basari-Backend/app/user"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	IdUser       int
	User         user.User `gorm:"foreignKey:IdUser"`
	IdBook       int
	Book         book.Book `gorm:"foreignKey:IdBook"`
	BorrowedDate time.Time
	DueDate      time.Time
	ReturnDate   time.Time
}
