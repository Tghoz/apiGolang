package controllers

import (
	"log"
	"net/http"

	dto "github.com/Tghoz/apiGolang/Dto"
	models "github.com/Tghoz/apiGolang/Model"
	repo "github.com/Tghoz/apiGolang/Repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostClient(c *gin.Context) {
	client := models.Clients{}
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client.ID = uuid.New()
	query := repo.Create(client)
	if query != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create client"})
		return
	}
	log.Println("Client created successfully")
	c.JSON(http.StatusCreated, client)
}

func GetClient(c *gin.Context) {
	clients, err := repo.FindAll(models.Clients{}, "Services", "History")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dataClient := []dto.ClientDto{}
	for _, c := range clients {
		dataService := []dto.ClienAndServerDto{}
		for _, s := range c.Services {
			dataService = append(dataService, dto.ClienAndServerDto{
				ID:    s.ID.String(),
				Name:  s.Name,
				Price: s.Price,
			})
		}
		dataHistory := []dto.HistoryDto{}
		for _, h := range c.History {
			dataHistory = append(dataHistory, dto.HistoryDtoMap(h))
		}
		dataClient = append(dataClient, dto.ClientDto{
			ID:        c.ID.String(),
			Name:      c.Name,
			Telephone: c.Telephone,
			Status:    c.Status,
			Services:  dataService,
			History:   dataHistory,
		})
	}

	if len(dataClient) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No content data"})
		return
	}
	c.JSON(http.StatusOK, dataClient)
}

func GetClientByID(c *gin.Context) {

	clientID := c.Param("id")
	client, err := repo.FindById(clientID, models.Clients{}, "Services", "History")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}
	dataService := []dto.ClienAndServerDto{}
	for _, s := range client.Services {
		dataService = append(dataService, dto.ClientAndServicesDto(s))
	}
	dataHistory := []dto.HistoryDto{}
	for _, h := range client.History {
		dataHistory = append(dataHistory, dto.HistoryDtoMap(h))
	}
	dataCient := dto.ClientDto{
		ID:        client.ID.String(),
		Name:      client.Name,
		Telephone: client.Telephone,
		Status:    client.Status,
		Services:  dataService,
		History:   dataHistory,
	}
	c.JSON(http.StatusOK, dataCient)
}

func DeleteClient(c *gin.Context) {

	clientID := c.Param("id")
	query := repo.Delete(clientID, models.Clients{})
	if query != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete client"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted client"})
}

func UpdateClient(c *gin.Context) {
	clientID := c.Param("id")
	var client models.Clients
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := repo.Update(clientID, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Client updated successfully"})
}
