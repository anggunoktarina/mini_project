package model

import (
	"time"

	"gorm.io/gorm"
)

type Savings_Progress struct {
	gorm.Model
	Template_ID  uint      `json:"template_id"`
	ID           uint      `json:"id"`
	Amount_Saved float64   `json:"amount_saved"`
	Category     string    `json:"category"`
	Target_Date  time.Time `json:"target_date"`
}
