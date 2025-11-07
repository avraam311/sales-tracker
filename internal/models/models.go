package models

type SaleDTO struct {
	Item   string  `json:"item" validate:"required"`
	Income float64 `json:"income" validate:"required"`
}

type SaleDB struct {
	ID     uint    `json:"id" validate:"required"`
	Item   string  `json:"item" validate:"required"`
	Income float64 `json:"income" validate:"required"`
}
