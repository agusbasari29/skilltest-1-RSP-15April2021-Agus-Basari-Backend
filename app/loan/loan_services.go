package loan

import "time"

type Services interface {
	CreateLoan(req RequestLoan) (Loan, error)
	Return(req RequestLoan) (Loan, error)
	GetLoan(req RequestLoan) *Loan
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) CreateLoan(req RequestLoan) (Loan, error) {
	loan := Loan{}
	loan.IdBook = req.IdBook
	loan.IdUser = req.IdUser
	loan.BorrowedDate = time.Now()
	loan.DueDate = loan.BorrowedDate.AddDate(0, 0, 3)

	newLoan, err := s.repository.InsertLoan(loan)
	if err != nil {
		return newLoan, err
	}
	return newLoan, nil
}

func (s *services) Return(req RequestLoan) (Loan, error) {
	ret := Loan{}
	ret.ID = req.ID
	ret.ReturnDate = time.Now()

	loan, err := s.repository.Return(ret)
	if err != nil {
		return loan, err
	}
	return loan, nil
}

func (s *services) GetLoan(req RequestLoan) *Loan {
	loan := Loan{}
	loan.ID = req.ID
	get := s.repository.FindLoan(loan)
	if get != nil {
		return get
	}
	return nil
}

func CreatePayment(day int) {

}
