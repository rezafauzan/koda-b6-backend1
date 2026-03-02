package data

type Response struct {
	Success      bool
	Messages     string
	ResponseBody any
}

//	{
//	    "id": 0,
//	    "avatar": "https://i.pravatar.cc/400?img=54",
//	    "fullname": "Reza Fauzan Adhima",
//	    "email": "rezafauzan@gmail.com",
//	    "phone": "085183356072",
//	    "address": "rezafauzan@gmail.com",
//	    "password": "dGVzdDEyMzQ=",
//	    "role": "user",
//	    "cart": [],
//	    "historyOrders": []
//	}
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