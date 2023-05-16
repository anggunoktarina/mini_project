package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required,max=20"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=5"`
	Token    string `json:"token"`
	UserID   uint   `json:"user_id"`
	Kas      []Kas
	Savings  []Savings
}
