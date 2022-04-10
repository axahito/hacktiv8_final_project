package models

import "time"

type Social struct {
	Name      string `gorm:"not null" json:"name" form:"name"`
	SocialURL string `gorm:"not null" json:"social_url" form:"social_url"`
	UserID    int    `gorm:"not null" json:"user_id" form:"user_id"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
