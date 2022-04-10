package models

import "time"

type Photo struct {
	Title     string `gorm:"not null" json:"title" form:"title"`
	Caption   string `json:"caption" form:"caption"`
	PhotoURL  string `gorm:"not null" json:"photo_url" form:"photo_url"`
	UserID    int    `gorm:"not null" json:"user_id" form:"user_id"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
