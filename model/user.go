package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Token    string    `json:"token"`
	Kas      []Kas     `gorm:"foreignKey:UserID"`
	Savings  []Savings `gorm:"foreignKey:UserID"`
}
