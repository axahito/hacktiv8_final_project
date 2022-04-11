package dto

import (
	"final_project/models"
	"time"
)

type comment struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Message   string `gorm:"not null" json:"message" form:"message"`
	UserID    int    `json:"user_id"`
	User      user   `json:"user"`
	PhotoID   int    `json:"photo_id"`
	Photo     photo  `json:"photo"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func MapComment(source []models.Comment) []comment {
	var result []comment

	for _, v := range source {
		result = append(result, comment{
			ID:      v.ID,
			Message: v.Message,
			PhotoID: v.PhotoID,
			Photo: photo{
				ID:       v.Photo.ID,
				Title:    v.Photo.Title,
				Caption:  v.Photo.Caption,
				PhotoURL: v.Photo.PhotoURL,
				UserID:   v.Photo.UserID,
			},
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			UserID:    v.UserID,
			User: user{
				ID:       v.User.ID,
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		})
	}

	return result
}
