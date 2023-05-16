package model

import (
	"gorm.io/gorm"
)

type Income struct {
	gorm.Model
	KasID          uint    `json:"kas_id"`
	Nominal_Amount float64 `json:"nominal_amount"`
	Category       string  `json:"category"`
	Date_Income    string  `json:"date_income"`
	Balance        int     `json:"balance"`
	Kas            []Kas
}
