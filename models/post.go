package models

type Post struct {
	Image   string `json:"img"`
	Content string `json:"content"`
	Author  int    `json:"author"`
	Likes   int    `json:"likes"`
}
