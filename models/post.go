package models

import "time"

type Post struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"postdate"`
	Title       string    `json:"title"`
	Desc        string    `json:"desc"`
	Content     string    `json:"content"`
	UserID      uint      `json:"userid"` //foreign key
	CommunityID uint      `json:"communityid"`
	User        User      `json:"user":gorm:"foreignkey:UserID"`
	Community   Community `json:"community":gorm:"foreignkey:CommunityID"`
}
