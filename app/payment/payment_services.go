package payment

type Services interface {
	CreatePayment(req RequestFinePayment) (FinePayment, error)
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) CreatePayment(req RequestFinePayment) (FinePayment, error) {
	payment := FinePayment{}
	payment.Amount = req.Amount
	payment.Receipt = req.Receipt
	payment.IdLoan = req.IdLoan
	newPayment, err := s.repository.InsertPayment(payment)
	if err != nil {
		return payment, err
	}
	return newPayment, nil
}
