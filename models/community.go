package models

type Community struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	ImageURL string `json:"image"`
	Link     string `json:"link"`
}
