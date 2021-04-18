package payment

import "gorm.io/gorm"

type Repository interface {
	InsertPayment(payment FinePayment) (FinePayment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) InsertPayment(payment FinePayment) (FinePayment, error) {
	err := r.db.Create(&payment).Error
	if err != nil {
		return payment, err
	}
	return payment, nil
}
