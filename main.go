package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	Id       int
	Fullname string
	Email    string
	Password string
}

var users []User

func main() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"messages": "success",
			"users":    users,
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
			ctx.JSON(200, gin.H{
				"Success":  true,
				"messages": "User found!",
				"users":    users[userFoundId],
			})
		} else {
			ctx.JSON(400, gin.H{
				"Success":  false,
				"Messages": "User not found !",
			})
		}
	})

	r.POST("/users", func(ctx *gin.Context) {
		data := User{}
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			ctx.JSON(400, gin.H{
				"Success":  false,
				"Messages": "Failed to create users",
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
				users = append(users, data)
				ctx.JSON(200, gin.H{
					"Success":  true,
					"Messages": "Users created",
					"Users":    users,
				})
			} else {
				ctx.JSON(400, gin.H{
					"Success":  false,
					"Messages": "Email allready used !",
				})
			}
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
			ctx.JSON(200, gin.H{
				"Success":  true,
				"messages": "User deleted!",
				"users":    users,
			})
		} else {
			ctx.JSON(400, gin.H{
				"Success":  false,
				"Messages": "User not found !",
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
				ctx.JSON(400, gin.H{
					"Success":  false,
					"Messages": "Failed to edit user!",
				})
			} else {
				if newData.Email != "" {
					users[userFoundId].Email = newData.Email
				}
				if newData.Password != "" {
					users[userFoundId].Password = newData.Password
				}
				if newData.Fullname != "" {
					users[userFoundId].Fullname = newData.Fullname
				}
				newId, err := strconv.Atoi(id)
				if err != nil {
					ctx.JSON(400, gin.H{
						"Success":  false,
						"Messages": "Edit user failed !",
					})
				} else {
					users[userFoundId].Id = newId
					ctx.JSON(200, gin.H{
						"Success":  true,
						"Messages": "Users edited",
						"Users":    users[userFoundId],
					})
				}
			}
		} else {
			ctx.JSON(400, gin.H{
				"Success":  false,
				"Messages": "User not found !",
			})
		}
	})

	r.Run("localhost:8888")
}
