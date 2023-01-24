package models

import "time"

type Comment struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"postdate"`
	Content   string    `json:"content"`
	PostID    int       `json:"postid"`
	UserID    int       `json:"userid"`
	Creator   string    `json:"creator"`
	Post      Post      `json:"post":gorm:"foreignkey:PostID"`
	User      User      `json:"user":gorm:"foreignkey:UserID"`
}
