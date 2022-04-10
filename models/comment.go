package models

import "time"

type Comment struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Message   string `gorm:"not null" json:"message" form:"message"`
	UserID    int    `gorm:"not null" json:"user_id" form:"user_id"`
	PhotoID   int    `gorm:"not null" json:"photo_id" form:"photo_id"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
