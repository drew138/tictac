package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Name     string `json:"name" gorm:"not null"`
	Surname  string `json:"surname" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	IsAdmin  bool   `json:"isAdmin" gorm:"default:false"`
}
