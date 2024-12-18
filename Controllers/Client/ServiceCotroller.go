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

	if len(dataService) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No content data"})
		return
	}

	c.JSON(http.StatusOK, dataService)
}

func GetServiceByID(c *gin.Context) {
	id := c.Param("id")
	serviceID, err := uuid.Parse(id)
	service := models.Services{}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	query := dataBase.Db.Preload("Clients").Where("id = ?", serviceID).First(&service, serviceID)
	if query.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	dataClient := []dto.ServiceAndClientDto{}
	for _, c := range service.Clients {
		dataClient = append(dataClient, dto.ServiceAndClientDto{
			ID:        c.ID.String(),
			Name:      c.Name,
			Telephone: c.Telephone,
			Status:    c.Status,
		})
	}
	dataService := dto.ServiceDto{
		ID:      service.ID.String(),
		Name:    service.Name,
		Price:   service.Price,
		Clients: dataClient,
	}

	c.JSON(http.StatusOK, dataService)
}
