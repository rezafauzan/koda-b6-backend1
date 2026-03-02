package main

import (
	"github.com/gin-gonic/gin"
	"rezafauzan/koda-b6-backend1/lib"
	"rezafauzan/koda-b6-backend1/middleware"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.GET("/users", lib.GetAllUsers)
	r.GET("/users/:id", lib.GetUserById)
	r.POST("/users", lib.Register)
	r.OPTIONS("/login", lib.Login)
	r.POST("/login", lib.Login)
	r.DELETE("/users/:id", lib.DeleteUserById)
	r.PATCH("/users/:id", lib.UpdateUserById)

	r.Run("localhost:8888")
}
