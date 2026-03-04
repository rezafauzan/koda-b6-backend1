package main

import (
	"fmt"
	"os"
	"rezafauzan/koda-b6-backend1/lib"
	"rezafauzan/koda-b6-backend1/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.GET("/users", lib.GetAllUsers)
	r.GET("/users/:id", lib.GetUserById)
	r.POST("/users", lib.Register)
	r.OPTIONS("/login", lib.Login)
	r.POST("/login", lib.Login)
	r.DELETE("/users/:id", lib.DeleteUserById)
	r.PATCH("/users/:id", lib.UpdateUserById)
	r.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
}
