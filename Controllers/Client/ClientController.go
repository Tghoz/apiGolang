package controllers

import (
	"net/http"

	dataBase "github.com/Tghoz/apiGolang/DataBase"
	dto "github.com/Tghoz/apiGolang/Dto"
	models "github.com/Tghoz/apiGolang/Model"
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
	if err := dataBase.Db.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create client"})
		return
	}
	c.JSON(http.StatusCreated, client)
}

func GetClient(c *gin.Context) {
	clients := []models.Clients{}
	if err := dataBase.Db.Preload("Services").Preload("History").Find(&clients).Error; err != nil {
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
	c.JSON(http.StatusOK, dataClient)
}

func GetClientByID(c *gin.Context) {
	id := c.Param("id")
	clientID, err := uuid.Parse(id)
	client := models.Clients{}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	query := dataBase.Db.Preload("Services").Preload("History").Where("id = ?", clientID).First(&client, clientID)
	if query.Error != nil {
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
