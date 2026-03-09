package routers

import (
	"rezafauzan/koda-b6-backend1/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewSwaggerRouters(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	docsRouter := router.Group("/docs")
	{
		docsRouter.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
