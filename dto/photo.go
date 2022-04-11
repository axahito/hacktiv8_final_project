package dto

import (
	"final_project/models"
	"time"
)

type photo struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Title     string `gorm:"not null" json:"title" form:"title"`
	Caption   string `json:"caption" form:"caption"`
	PhotoURL  string `gorm:"not null" json:"photo_url" form:"photo_url"`
	UserID    int    `json:"user_id"`
	User      user   `json:"user"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func MapPhoto(source []models.Photo) []photo {
	var result []photo

	for _, v := range source {
		result = append(result, photo{
			ID:        v.ID,
			Title:     v.Title,
			Caption:   v.Caption,
			PhotoURL:  v.PhotoURL,
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
