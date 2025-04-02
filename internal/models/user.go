package models

import "time"

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Username       string `gorm:"uniqueIndex;not null"`
	Email          string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	ProfilePicture string `gorm:"default:null"`
	Bio            string
	FollowersCount uint
	FollowingCount uint
	PostsCount     uint
	CreatedAt      time.Time

	// Relationships
	Followers []User `gorm:"many2many:user_followers;"`
	Following []User `gorm:"many2many:user_followers;"`
	Posts     []Post `gorm:"foreignKey:UserID"`
	Likes     []Like `gorm:"foreignKey:UserID"`
}
