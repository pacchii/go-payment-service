package service

import (
	"go-payment-service/internal/domain"
	"go-payment-service/internal/repository"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type PaymentService interface {
	CreatePayment(amount float64) domain.Payment
	GetPayment(id string) (domain.Payment, bool)
	FindAll() ([]domain.Payment, int)
}

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) CreatePayment(amount float64) domain.Payment {
	p := domain.Payment{
		Id:     uuid.New().String(),
		Amount: amount,
		Status: domain.Pending,
	}
	s.repo.Save(p)

	go s.processPayment(p.Id)

	return p

}

func (s *paymentService) processPayment(id string) {

	time.Sleep(2 * time.Second)

	p, ok := s.repo.FindById(id)

	if !ok {
		return
	}

	if rand.Intn(2) == 0 {
		p.Status = domain.Success
	} else {
		p.Status = domain.Failure
	}

	s.repo.Save(p)
}

func (s *paymentService) GetPayment(id string) (domain.Payment, bool) {
	return s.repo.FindById(id)
}

func (s *paymentService) FindAll() ([]domain.Payment, int) {
	return s.repo.FindAll()
}
