package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"strconv"
	"strings"
)

type Response struct {
	Success      bool
	Messages     string
	ResponseBody any
}
// {
//     "id": 0,
//     "avatar": "https://i.pravatar.cc/400?img=54",
//     "fullname": "Reza Fauzan Adhima",
//     "email": "rezafauzan@gmail.com",
//     "phone": "085183356072",
//     "address": "rezafauzan@gmail.com",
//     "password": "dGVzdDEyMzQ=",
//     "role": "user",
//     "cart": [],
//     "historyOrders": []
// }
type User struct {
	Id       int
	Avatar string
	Firstname string
	Lastname string
	Email    string
	Phone string
	Address string
	Password string
	Role string
}

var users []User
var loggedInUser User

func main() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success:      true,
			Messages:     "Success",
			ResponseBody: users,
		})
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		userFoundId := -1
		for x := range users {
			if strconv.Itoa(users[x].Id) == id {
				userFoundId = x
			}
		}
		if userFoundId > -1 {
			ctx.JSON(200, Response{
				Success:      true,
				Messages:     "User found!",
				ResponseBody: users[userFoundId],
			})
		} else {
			ctx.JSON(400, Response{
				Success:      false,
				Messages:     "User not found !",
				ResponseBody: "",
			})
		}
	})

	r.POST("/users", func(ctx *gin.Context) {
		data := User{}
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			ctx.JSON(400, Response{
				Success:      false,
				Messages:     "Failed to create users",
				ResponseBody: "",
			})
		} else {
			emailExist := 0
			for x := range users {
				if users[x].Email == data.Email {
					emailExist++
				}
			}
			if emailExist == 0 {
				data.Id = len(users)
				if len(data.Firstname) < 4 {
					ctx.JSON(400, Response{
						Success:      false,
						Messages:     "Firstname minimal 4 characters!",
						ResponseBody: "",
					})
					return
				}
				
				if len(data.Lastname) < 4 {
					ctx.JSON(400, Response{
						Success:      false,
						Messages:     "Lastname minimal 4 characters!",
						ResponseBody: "",
					})
					return
				}

				if len(data.Email) < 4 || strings.Contains(data.Email, "@") != true {
					ctx.JSON(400, Response{
						Success:      false,
						Messages:     "Email minimal 4 characters and must be a valid email!",
						ResponseBody: "",
					})
					return
				}

				if len(data.Phone) < 10 {
					ctx.JSON(400, Response{
						Success:      false,
						Messages:     "Phone numbers minimal 10 digits",
						ResponseBody: "",
					})
					return
				}

				if len(data.Address) < 10 {
					ctx.JSON(400, Response{
						Success:      false,
						Messages:     "Address minimal 10 characters",
						ResponseBody: "",
					})
					return
				}

				if len(data.Password) < 8 {
					ctx.JSON(400, Response{
						Success:      false,
						Messages:     "Password too weak minimal 8 characters!",
						ResponseBody: "",
					})
					return
				} else {
					argon := argon2.DefaultConfig()
					hash, err := argon.HashEncoded([]byte(data.Password))
					if err != nil {
						ctx.JSON(400, Response{
							Success:      false,
							Messages:     "System fail to proses password!",
							ResponseBody: "",
						})
						return
					} else {
						data.Password = string(hash)
					}
				}
				data.Avatar = "https://i.pravatar.cc/400?img=54"
				data.Role = "Member"
				users = append(users, data)
				ctx.JSON(200, Response{
					Success:      true,
					Messages:     "Users created",
					ResponseBody: users,
				})
			} else {
				ctx.JSON(400, Response{
					Success:      false,
					Messages:     "Email allready used !",
					ResponseBody: "",
				})
			}
		}
	})

	r.POST("/login", func(ctx *gin.Context) {
		data := User{}
		err := ctx.ShouldBindJSON(&data)
		if len(loggedInUser.Email) < 1 {
			if err != nil {
				ctx.JSON(400, Response{
					Success:      false,
					Messages:     "Login failed",
					ResponseBody: "",
				})
				return
			} else {
				for x := range users {
					if users[x].Email == data.Email {
						correct, err := argon2.VerifyEncoded([]byte(data.Password), []byte(users[x].Password))
						if err != nil {
							ctx.JSON(400, Response{
								Success:      false,
								Messages:     "Email or password wrong!",
								ResponseBody: "",
							})
							return
						} else if correct {
							loggedInUser = users[x]
							loggedInUser.Password = "Hidden"
							ctx.JSON(200, Response{
								Success:      false,
								Messages:     "Login success! wellcome back " + users[x].Firstname + " " + users[x].Lastname,
								ResponseBody: loggedInUser,
							})
							return
						}
					} else {
						ctx.JSON(400, Response{
							Success:      false,
							Messages:     "Email or password wrong!",
							ResponseBody: "",
						})
						return
					}
				}
			}
		} else {
			ctx.JSON(200, Response{
				Success:      false,
				Messages:     "You allready logged in!!",
				ResponseBody: loggedInUser,
			})
			return
		}
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		userFoundId := -1
		for x := range users {
			if strconv.Itoa(users[x].Id) == id {
				userFoundId = x
			}
		}
		if userFoundId > -1 {
			users = append(users[:userFoundId], users[userFoundId+1:]...)
			ctx.JSON(200, Response{
				Success:      true,
				Messages:     "User deleted!",
				ResponseBody: users,
			})
		} else {
			ctx.JSON(400, Response{
				Success:      false,
				Messages:     "User not found !",
				ResponseBody: "",
			})
		}
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		userFoundId := -1

		for x := range users {
			if strconv.Itoa(users[x].Id) == id {
				userFoundId = x
			}
		}

		if userFoundId > -1 {
			newData := User{}
			err := ctx.ShouldBindJSON(&newData)
			if err != nil {
				ctx.JSON(400, Response{
					Success:      false,
					Messages:     "Failed to edit user!",
					ResponseBody: "",
				})
				return
			} else {
				if newData.Email != "" {
					emailExist := 0
					for x := range users {
						if users[x].Email == newData.Email {
							emailExist++
						}
					}
					if emailExist == 0 {
						users[userFoundId].Email = newData.Email
					} else {
						ctx.JSON(400, Response{
							Success:      false,
							Messages:     "Email allready used !",
							ResponseBody: "",
						})
						return
					}
				}
				if newData.Password != "" {
					argon := argon2.DefaultConfig()
					hash, err := argon.HashEncoded([]byte(newData.Password))
					if err != nil {
						ctx.JSON(400, Response{
							Success:      false,
							Messages:     "System fail to proses password!",
							ResponseBody: "",
						})
						return
					} else {
						newData.Password = string(hash)
						users[userFoundId].Password = newData.Password
					}
				}

				if newData.Firstname != "" {
					users[userFoundId].Firstname = newData.Firstname
				}

				if newData.Lastname != "" {
					users[userFoundId].Lastname = newData.Lastname
				}
				
				newId, err := strconv.Atoi(id)
				if err != nil {
					ctx.JSON(400, Response{
						Success:      false,
						Messages:     "Edit user failed !",
						ResponseBody: "",
					})
					return
				} else {
					users[userFoundId].Id = newId
					ctx.JSON(200, Response{
						Success:      true,
						Messages:     "Users edited",
						ResponseBody: users[userFoundId],
					})
				}
			}
		} else {
			ctx.JSON(400, Response{
				Success:      false,
				Messages:     "User not found !",
				ResponseBody: "",
			})
			return
		}
	})

	r.Run("localhost:8888")
}
