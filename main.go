package main

import "github.com/gin-gonic/gin"

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

	r.POST("/users", func(ctx *gin.Context) {
		data := User{}
		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			ctx.JSON(400, gin.H{
				"Success":  false,
				"Messages": "Failed to create users",
			})
		} else {
			ctx.JSON(200, gin.H{
				"Success":  true,
				"Messages": "Users created",
				"User": data,
			})
		}
	})

	r.Run("localhost:8888")
}
