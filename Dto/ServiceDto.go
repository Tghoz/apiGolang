package dto

import models "github.com/Tghoz/apiGolang/Model"

type serviceDto struct {
	ID     string
	Name   string
	Price  float64
	Client []ClientDto
}

func ServiceDtoMap(s models.Services) serviceDto {

	return serviceDto{
		ID:     s.ID.String(),
		Name:   s.Name,
		Price:  s.Price,
		Client: []ClientDto{},
	}
	
}
