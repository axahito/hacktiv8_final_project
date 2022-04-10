package models

import "time"

type Comment struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Message   string `gorm:"not null" json:"message" form:"message"`
	UserID    int    `gorm:"not null"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PhotoID   int    `gorm:"not null" json:"photo_id" form:"photo_id"`
	Photo     Photo  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
