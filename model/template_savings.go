package model

import (
	"time"

	"gorm.io/gorm"
)

type Template_Savings struct {
	gorm.Model
	Template_Name          string    `json:"template_name"`
	Time_Period            time.Time `json:"time_period"`
	Additional_Information string    `json:"additional_information"`
}
