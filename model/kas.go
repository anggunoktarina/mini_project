package model

import "gorm.io/gorm"

type Kas struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	IncomeID  uint `json:"income_id"`
	ExpenseID uint `json:"expense_id"`
	Total_Kas int  `json:"total_kas"`
	Balance   int  `json:"balance"`
}
