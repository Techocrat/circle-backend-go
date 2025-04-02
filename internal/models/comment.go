package models

import "time"

type Comment struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      
    PostID    uint      
    CommentText string 
    CreatedAt time.Time 

    // Relationships
    User      User      `gorm:"foreignKey:UserID"` 
}
