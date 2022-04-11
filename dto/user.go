package dto

type user struct {
	ID       int    `json:"-"`
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email"`
	// Age      int    `gorm:"not null" json:"age" form:"age"`
}
