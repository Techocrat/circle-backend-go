package models

import "time"

type UserFollower struct {
	ID         uint `gorm:"primaryKey"`
	FollowerID uint `gorm:"not null"`
	FollowedID uint `gorm:"not null"`
	CreatedAt  time.Time
}
