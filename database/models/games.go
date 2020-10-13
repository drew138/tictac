package models

import "time"

// Game - Information pertaining searches done by each user
type Game struct {
	ID         uint `gorm:"primaryKey"`
	PlayersOne User `gorm:"foreignKey:ID;constraint:OnDelete:SET NULL;"`
	PlayersTwo User `gorm:"foreignKey:ID;constraint:OnDelete:SET NULL;"`
	CreatedAt  time.Time
}
