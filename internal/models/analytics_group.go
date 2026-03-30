package models

import "time"

type AnalyticsGroupDB struct {
	Period       time.Time `json:"period" db:"period"`
	Sum          float64   `json:"sum" db:"sum"`
	Avg          float64   `json:"avg" db:"avg"`
	Count        int64     `json:"count" db:"count"`
	Median       float64   `json:"median" db:"median"`
	Percentile90 float64   `json:"percentile_90" db:"percentile_90"`
}
