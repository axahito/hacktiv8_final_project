package models

import "time"

type Social struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null" json:"name" form:"name"`
	SocialURL string `gorm:"not null" json:"social_url" form:"social_url"`
	UserID    int    `gorm:"not null"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedAt time.Time
	CreatedAt time.Time
}
