package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserRouters(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"OK": "OK",
			})
		})
		// userRoutes.GET("/users", handler.GetAllUsers)
		// userRoutes.GET("/users/:id", handler.GetUserById)
		// userRoutes.POST("/users", handler.Register)
		// userRoutes.POST("/login", handler.Login)
		// userRoutes.DELETE("/users/:id", handler.DeleteUserById)
		// userRoutes.PATCH("/users/:id", handler.UpdateUserById)
	}
}
