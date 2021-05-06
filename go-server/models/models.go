package models

// User models
type User struct {
	ID       int    `json:"id" sql:"AUTO_INCREMENT"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// UserResponse models
type UserResponse struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
