package repository

import (
	"go-payment-service/internal/domain"
	"sync"
)

type PaymentRepository interface {
	Save(payment domain.Payment)
	FindById(id string) (domain.Payment, bool)
	FindAll() ([]domain.Payment, int)
}

type PaymentRepo struct {
	data map[string]domain.Payment
	mu   sync.RWMutex
}

func NewPaymentRepo() *PaymentRepo {
	return &PaymentRepo{
		data: make(map[string]domain.Payment),
	}
}

func (p *PaymentRepo) Save(payment domain.Payment) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.data[payment.Id] = payment
}

func (r *PaymentRepo) FindById(id string) (domain.Payment, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	p, ok := r.data[id]
	return p, ok
}

func (r *PaymentRepo) FindAll() ([]domain.Payment, int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var payments []domain.Payment
	for _, p := range r.data {
		payments = append(payments, p)
	}
	return payments, len(r.data)
}
