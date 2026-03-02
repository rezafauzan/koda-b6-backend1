package lib

import (
	"github.com/gin-gonic/gin"
	"rezafauzan/koda-b6-backend1/data"
	"strconv"
)

func DeleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	userFoundId := -1
	for x := range data.Users {
		if strconv.Itoa(data.Users[x].Id) == id {
			userFoundId = x
		}
	}
	if userFoundId > -1 {
		data.Users = append(data.Users[:userFoundId], data.Users[userFoundId+1:]...)
		ctx.JSON(200, data.Response{
			Success:      true,
			Messages:     "User deleted!",
			ResponseBody: data.Users,
		})
	} else {
		ctx.JSON(400, data.Response{
			Success:      false,
			Messages:     "User not found !",
			ResponseBody: "",
		})
	}
}
