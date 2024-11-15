package dto

import models "github.com/Tghoz/apiGolang/Model"

type ServiceDto struct {
	ID      string
	Name    string
	Price   float64
	Clients []ServiceAndClientDto
}

type ClienAndServerDto struct {
	ID    string
	Name  string
	Price float64
}

func ClientAndServicesDto(s models.Services) ClienAndServerDto {
	return ClienAndServerDto{
		ID:    s.ID.String(),
		Name:  s.Name,
		Price: s.Price,
	}
}
