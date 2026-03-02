package lib

import (
	"github.com/gin-gonic/gin"
	"rezafauzan/koda-b6-backend1/data"
)

func GetAllUsers(ctx *gin.Context) {
	ctx.JSON(200, data.Response{
		Success:      true,
		Messages:     "Success",
		ResponseBody: data.Users,
	})
}
