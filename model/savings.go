package model

import (
	"gorm.io/gorm"
)

type Savings struct {
	gorm.Model
	UserID                 uint    `json:"user_id"`
	Amount_Saved           float64 `json:"amount_saved"`
	Category               string  `json:"category"`
	Target_Date            string  `json:"target_date"`
	Start_Date             string  `json:"start_date"`
	Additional_Information string  `json:"additional_information"`
	Total_Saved            float64 `json:"total_saved"`
}
