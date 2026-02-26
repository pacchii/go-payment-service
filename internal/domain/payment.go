package domain

type PaymentStatus string

const (
	Pending PaymentStatus = "PENDING"
	Success PaymentStatus = "SUCCESS"
	Failure PaymentStatus = "FAILURE"
)

type Payment struct {
	Id     string        `json:"id"`
	Amount float64       `json:"amount"`
	Status PaymentStatus `json:"status"`
}
