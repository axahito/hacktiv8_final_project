package dto

import "time"

type Social struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null" json:"name" form:"name"`
	SocialURL string `gorm:"not null" json:"social_url" form:"social_url"`
	UserID    int    `json:"user_id"`
	User      user   `json:"user"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
