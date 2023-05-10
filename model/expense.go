package model

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Kas_ID         uint      `json:"kas_id"`
	Nominal_Amount float64   `json:"nominal_amount"`
	Category       string    `json:"category"`
	Date_Expense   time.Time `json:"date_expense"`
	Kas            []Kas
	Balance        int `json:"balance"`
}
