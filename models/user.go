package models

import (
	"final_project/helpers"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"not null;uniqueIndex" json:"username" form:"username"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email" form:"email"`
	Password  string    `gorm:"not null" json:"password" form:"password"`
	Age       int       `gorm:"not null" json:"age" form:"age"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPass(u.Password)

	return
}
