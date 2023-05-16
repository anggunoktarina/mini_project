package model

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	KasID          uint    `json:"kas_id"`
	Nominal_Amount float64 `json:"nominal_amount"`
	Category       string  `json:"category"`
	Date_Expense   string  `json:"date_expense"`
	Kas            []Kas
	Balance        int `json:"balance"`
}
