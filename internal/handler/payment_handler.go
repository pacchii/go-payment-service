package handler

import (
	"go-payment-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	payment := h.service.CreatePayment(req.Amount)

	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
	id := c.Param("id")

	p, ok := h.service.GetPayment(id)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payment Not Found"})
	}
	c.JSON(http.StatusOK, p)
}

func (h *PaymentHandler) AllPayments(c *gin.Context) {

	p, _ := h.service.FindAll()

	c.JSON(http.StatusOK, p)
}
