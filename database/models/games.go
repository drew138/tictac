package models

import "time"

// Game - Information pertaining searches done by each user
type Game struct {
	ID         int  `gorm:"primaryKey"`
	PlayersOne User `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PlayersTwo User `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt  time.Time
}
