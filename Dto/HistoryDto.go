package dto

import (
	models "github.com/Tghoz/apiGolang/Model"
)

type HistoryDto struct {
	ID       string
	ClientID string
	Amount   float64
	Type     string
	Date     string
}

func HistoryDtoMap(h models.Payments) HistoryDto {
	return HistoryDto{
		ID:       h.ID.String(),
		ClientID: h.ClientID.String(),
		Amount:   h.Amount,
		Type:     h.Type,
		Date:     h.Date.String(),
	}
}
