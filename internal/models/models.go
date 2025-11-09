package models

type SaleDTO struct {
	Item   string  `json:"item" validate:"required" db:"item"`
	Income float64 `json:"income" validate:"required" db:"income"`
}

type SaleDB struct {
	ID        uint      `json:"id" validate:"required" db:"id"`
	Item      string    `json:"item" validate:"required" db:"item"`
	Income    float64   `json:"income" validate:"required" db:"income"`
}

type AnalyticsDB struct {
	Sum          float64 `json:"sum" validate:"required" db:"sum"`
	Avg          float64 `json:"avg" validate:"required" db:"avg"`
	Count        int64   `json:"count" validate:"required" db:"count"`
	Median       float64 `json:"median" validate:"required" db:"median"`
	Percentile90 float64 `json:"percentile_90" validate:"required" db:"percentile_90"`
}
