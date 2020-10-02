package models

import "time"

// Game - Information pertaining searches done by each user
type Game struct {
	ID        int  `gorm:"primaryKey"`
	PlayerOne User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PlayerTwo User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
}
