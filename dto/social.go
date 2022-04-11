package dto

import (
	"final_project/models"
	"time"
)

type social struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null" json:"name" form:"name"`
	SocialURL string `gorm:"not null" json:"social_url" form:"social_url"`
	UserID    int    `json:"user_id"`
	User      user   `json:"user"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func MapSocial(source []models.Social) []social {
	var result []social

	for _, v := range source {
		result = append(result, social{
			ID:        v.ID,
			Name:      v.Name,
			SocialURL: v.SocialURL,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			UserID:    v.UserID,
			User: user{
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		})
	}

	return result
}
