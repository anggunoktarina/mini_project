package model

import (
	"time"

	"gorm.io/gorm"
)

type Income struct {
	gorm.Model
	Kas_ID         uint      `json:"kas_id"`
	Nominal_Amount float64   `json:"nominal_amount"`
	Category       string    `json:"category"`
	Date_Income    time.Time `json:"date_income"`
	Balance        int       `json:"balance"`
	Kas            []Kas
}
