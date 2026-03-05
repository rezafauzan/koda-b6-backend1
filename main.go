package main

import (
	"fmt"
	"os"
	"rezafauzan/koda-b6-backend1/lib"
	"rezafauzan/koda-b6-backend1/docs"
	"rezafauzan/koda-b6-backend1/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files" 
)

// @title Koda-B6-Backend1
// @version 1.0
// @description Backend app built with gin swagger
// @host localhost:8888
// @BasePath /
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
	docs.SwaggerInfo.BasePath = "/"
	docsRouter := r.Group("/docs")
	{
		docsRouter.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	r.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
}
