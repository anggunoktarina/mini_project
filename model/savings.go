package model

import (
	"time"

	"gorm.io/gorm"
)

type Savings struct {
	gorm.Model
	Amount_Saved           float64   `json:"amount_saved"`
	Category               string    `json:"category"`
	Target_Date            time.Time `json:"target_date"`
	Start_Date             time.Time `json:"start_date"`
	Additional_Information string    `json:"additional_information"`
	Total_Saved            float64   `json:"total_saved"`
	User_ID                int       `json:"user_id"`
}
