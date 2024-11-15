package controllers

import (
	"net/http"

	dataBase "github.com/Tghoz/apiGolang/DataBase"
	dto "github.com/Tghoz/apiGolang/Dto"
	models "github.com/Tghoz/apiGolang/Model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostService(c *gin.Context) {
	service := models.Services{}
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.ID = uuid.New()
	if err := dataBase.Db.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create service"})
		return
	}
	c.JSON(http.StatusCreated, service)
}

func GetService(c *gin.Context) {
	services := []models.Services{}
	if err := dataBase.Db.Preload("Clients").Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dataService := []dto.ServiceDto{}
	for _, s := range services {
		dataClient := []dto.ServiceAndClientDto{}
		for _, c := range s.Clients {
			dataClient = append(dataClient, dto.ServiceAndClientDto{
				ID:        c.ID.String(),
				Name:      c.Name,
				Telephone: c.Telephone,
				Status:    c.Status,
			})
		}

		dataService = append(dataService, dto.ServiceDto{
			ID:      s.ID.String(),
			Name:    s.Name,
			Price:   s.Price,
			Clients: dataClient,
		})
	}

	c.JSON(http.StatusOK, dataService)
}
