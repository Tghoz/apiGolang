package dto

import (
	models "github.com/Tghoz/apiGolang/Model"
)

type HistoryDto struct {
	ID     string
	Amount float64
	Type   string
	Date   string
}

func HistoryDtoMap(h models.Payments) HistoryDto {
	return HistoryDto{
		ID:     h.ID.String(),
		Amount: h.Amount,
		Type:   h.Type,
		Date:   h.Date.String(),
	}
}
