package models

import "time"

type Post struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint
    Description string
    ImageURL    string
    CreatedAt   time.Time
}
