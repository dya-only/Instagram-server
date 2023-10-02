package models

type User struct {
	Email     string `json:"email"`
	UserName  string `json:"username"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Bookmarks string `json:"bookmarks"`
	Likes     string `json:"likes"`
	Follower  string `json:"follower"`
	Following string `json:"following"`
}
