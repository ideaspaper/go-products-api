package model

import "productsapi/internal/dto/response"

type Product struct {
	ID          int
	Name        string
	Description string
	Quantity    int
}

func (p Product) ToDTO() *response.Product {
	return &response.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Quantity:    p.Quantity,
	}
}
