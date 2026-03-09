package main

import (
	"fmt"
	"os"
	"rezafauzan/koda-b6-backend1/internal/routers"
	"rezafauzan/koda-b6-backend1/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Koda-B6-Backend1
// @version 1.0
// @description Backend app built with gin swagger
// @host localhost:8888
// @BasePath /
func main() {
	godotenv.Load()

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	routers.NewUserRouters(router)
	routers.NewSwaggerRouters(router)

	router.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
}
