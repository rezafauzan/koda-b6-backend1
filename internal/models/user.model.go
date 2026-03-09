package models

type User struct {
	Id        int `json:"id"`
	Avatar    string `json:"avatar"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type CreateUserRequest struct{
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Password  string `json:"password"`
}