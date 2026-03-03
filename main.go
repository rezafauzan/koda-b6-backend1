package main

import (
	"context"
	"fmt"
	"os"
	"rezafauzan/koda-b6-backend1/lib"
	"rezafauzan/koda-b6-backend1/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env!")
		return
	}

	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		fmt.Println("Failed parse database config!")
		return
	}

	connection, err := pgx.Connect(context.Background(), connConfig.ConnString())
	if err != nil {
		fmt.Println("Could not connect to database!")
		return
		}else{
		fmt.Println("Connection to database established!")
	}
	connection.Close(context.Background())

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
