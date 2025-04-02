package models

import "time"

type Post struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	Description   string
	ImageURL      string
	LikesCount    uint
	CommentsCount uint
	CreatedAt     time.Time
	UserLiked     bool

	// Relationships
	Likes    []Like    `gorm:"foreignKey:PostID"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}
