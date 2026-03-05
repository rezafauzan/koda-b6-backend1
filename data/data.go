package data

type Response struct {
	Success      bool
	Messages     string
	ResponseBody any
}

type UserStruct struct {
	Id        int
	Avatar    string
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	Address   string
	Password  string
	Role      string
}

var Users []UserStruct
var User UserStruct