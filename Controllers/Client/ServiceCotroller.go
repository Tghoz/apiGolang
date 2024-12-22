package controllers

import (
	"net/http"

	dto "github.com/Tghoz/apiGolang/Dto"
	models "github.com/Tghoz/apiGolang/Model"
	repo "github.com/Tghoz/apiGolang/Repository"
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
	query := repo.Create(service)
	if query != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create service"})
		return
	}
	c.JSON(http.StatusCreated, service)
}

func GetService(c *gin.Context) {
	services, err := repo.FindAll(models.Services{}, "Clients")
	if err != nil {
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
	service, err := repo.FindById(id, models.Services{}, "Clients")
	if err != nil {
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

func Delete(c *gin.Context) {
	id := c.Param("id")
	query := repo.Delete(id, models.Services{})
	if query != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete service"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted"})
}

func UpdateService(c *gin.Context) {
	id := c.Param("id")
	service := models.Services{}
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := repo.Update(id, service)
	if query != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update service"})
		return
	}
	c.JSON(http.StatusOK, service)

}
