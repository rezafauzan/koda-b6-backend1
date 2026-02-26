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
	r.Run("localhost:8888")
}
