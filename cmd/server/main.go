package main

import (
	"go-payment-service/internal/handler"
	"go-payment-service/internal/repository"
	"go-payment-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo := repository.NewPaymentRepo()
	svc := service.NewPaymentService(repo)
	h := handler.NewPaymentHandler(svc)

	r.POST("/payment", h.CreatePayment)
	r.GET("/payment/:id", h.GetPayment)
	r.GET("/payments", h.AllPayments)

	r.Run(":9090")

}
