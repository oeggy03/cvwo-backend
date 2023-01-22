package models

type Comment struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	PostID  int    `json:"postid"`
	UserID  string `json:"userid"`
	Post    Post   `json:"post":gorm:"foreignkey:PostID"`
	User    User   `json:"user":gorm:"foreignkey:UserID"`
}
