package model

import "gorm.io/gorm"

type Kas struct {
	gorm.Model
	User_ID    int `json:"user_id"`
	Income_ID  int `json:"income_id"`
	Expense_ID int `json:"expense_id"`
	Total_Kas  int `json:"total_kas"`
	Balance    int `json:"balance"`
}
