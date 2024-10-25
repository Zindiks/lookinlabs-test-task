package model

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email`
}

func CreateUser() *User {
	return &User{}
}
