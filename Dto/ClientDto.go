package dto

import (
	models "github.com/Tghoz/apiGolang/Model"
)

type ClientDto struct {
	ID        string
	Name      string
	Telephone string
	Status    string
	Services  []ClienAndServerDto
	History   []HistoryDto
}

type ServiceAndClientDto struct {
	ID        string
	Name      string
	Telephone string
	Status    string
}


func ClientDtoMap(c models.Clients) ClientDto {
	return ClientDto{
		ID:        c.ID.String(),
		Name:      c.Name,
		Telephone: c.Telephone,
		Status:    c.Status,
		Services:  []ClienAndServerDto{},
		History:   []HistoryDto{},
	}
}



