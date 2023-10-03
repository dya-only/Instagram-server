package models

type Comment struct {
	Author   int    `json:"author"`
	PostId   int    `json:"postid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Content  string `json:"content"`
	Likes    int    `json:"likes"`
}
