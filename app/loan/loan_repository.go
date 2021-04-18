package loan

import "gorm.io/gorm"

type Repository interface {
	InsertLoan(loan Loan) (Loan, error)
	FindLoan(loan Loan) *Loan
	Return(loan Loan) (Loan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertLoan(loan Loan) (Loan, error) {
	err := r.db.Create(&loan).Error
	if err != nil {
		return loan, err
	}
	return loan, nil
}

func (r *repository) FindLoan(loan Loan) *Loan {
	err := r.db.First(&loan).Error
	if err == nil {
		return &loan
	}
	return nil
}

func (r *repository) Return(loan Loan) (Loan, error) {
	err := r.db.Save(&loan).Error
	if err != nil {
		return loan, err
	}
	return loan, nil
}
