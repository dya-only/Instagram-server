package models

type Comment struct {
	Author  int    `json:"author"`
	PostId  int    `json:"postid"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
}
