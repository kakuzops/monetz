package models

import (
	"gorm.io/gorm"
)

type Seller struct {
	gorm.Model
	Name          string  `json:"name"`
	TaxPercentual float64 `json:"tax_percentual"`
	FixTax        float64 `json:"fix_tax"`
	UserID        string  `json:"user_id"`
}
