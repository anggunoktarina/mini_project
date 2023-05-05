package model

import "gorm.io/gorm"

type Kas struct {
	gorm.Model
	Kas_ID     int `json:"kas_id"`
	Income_ID  int `json:"income_id"`
	Expense_ID int `json:"expense_id"`
	Total      int `json:"total"`
}
