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

	clientDto := []dto.ClientDto{}
	// serviceDto := []dto.ServiceDto{}
	historyDto := []dto.HistoryDto{}

	for _, client := range clients {

		for _, h := range client.History {
			historyDto = append(historyDto, dto.HistoryDtoMap(h)) // Asegúrate de que esta función sea correcta
		}

		clientDto = append(clientDto,
			dto.ClientDto{
				ID:        client.ID.String(),
				Name:      client.Name,
				Telephone: client.Telephone,
				Status:    client.Status,
				// Services:  serviceDto,
				History: historyDto,
			})
	}
	c.JSON(http.StatusOK, clientDto)

}
