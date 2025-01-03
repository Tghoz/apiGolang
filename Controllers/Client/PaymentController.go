package controllers

import (
	"net/http"

	dto "github.com/Tghoz/apiGolang/Dto"
	models "github.com/Tghoz/apiGolang/Model"
	repo "github.com/Tghoz/apiGolang/Repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PaymentController is a struct that will hold the methods to handle the payment routes

func PostPayment(c *gin.Context) {
	payment := models.Payments{}
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment.ID = uuid.New()
	query := repo.Create(payment)
	if query != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create payment"})
		return
	}
	c.JSON(http.StatusCreated, payment)
}

func GetPayment(c *gin.Context) {
	payments, err := repo.FindAll(models.Payments{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	payment := []dto.HistoryDto{}
	for _, p := range payments {
		payment = append(payment, dto.HistoryDtoMap(p))
	}
	if len(payments) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No content data"})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func GetPaymentByID(c *gin.Context) {
	id := c.Param("id")
	query, err := repo.FindById(id, models.Payments{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	paymentDto := dto.HistoryDtoMap(*query)
	c.JSON(http.StatusOK, paymentDto)
}
