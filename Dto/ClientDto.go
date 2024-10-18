package dto

import (
	models "github.com/Tghoz/apiGolang/Model"
)

type ClientDto struct {
	ID        string
	Name      string
	Telephone string
	Status    string
	Services  []models.Services
	History   []HistoryDto
}
