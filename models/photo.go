package models

import "time"

type Photo struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Title     string `gorm:"not null" json:"title" form:"title"`
	Caption   string `json:"caption" form:"caption"`
	PhotoURL  string `gorm:"not null" json:"photo_url" form:"photo_url"`
	UserID    int    `gorm:"not null"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
