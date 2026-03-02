package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"rezafauzan/koda-b6-backend1/data"
	"strconv"
)

func UpdateUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	userFoundId := -1

	for x := range data.Users {
		if strconv.Itoa(data.Users[x].Id) == id {
			userFoundId = x
		}
	}

	if userFoundId > -1 {
		newData := data.UserStruct{}
		err := ctx.ShouldBindJSON(&newData)
		if err != nil {
			ctx.JSON(400, data.Response{
				Success:      false,
				Messages:     "Failed to edit user!",
				ResponseBody: "",
			})
			return
		} else {
			if newData.Email != "" {
				emailExist := 0
				for x := range data.Users {
					if data.Users[x].Email == newData.Email {
						emailExist++
					}
				}
				if emailExist == 0 {
					data.Users[userFoundId].Email = newData.Email
				} else {
					ctx.JSON(400, data.Response{
						Success:      false,
						Messages:     "Email allready used !",
						ResponseBody: "",
					})
					return
				}
			}
			if newData.Password != "" {
				argon := argon2.DefaultConfig()
				hash, err := argon.HashEncoded([]byte(newData.Password))
				if err != nil {
					ctx.JSON(400, data.Response{
						Success:      false,
						Messages:     "System fail to proses password!",
						ResponseBody: "",
					})
					return
				} else {
					newData.Password = string(hash)
					data.Users[userFoundId].Password = newData.Password
				}
			}

			if newData.Firstname != "" {
				data.Users[userFoundId].Firstname = newData.Firstname
			}

			if newData.Lastname != "" {
				data.Users[userFoundId].Lastname = newData.Lastname
			}

			newId, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(400, data.Response{
					Success:      false,
					Messages:     "Edit user failed !",
					ResponseBody: "",
				})
				return
			} else {
				data.Users[userFoundId].Id = newId
				ctx.JSON(200, data.Response{
					Success:      true,
					Messages:     "Users edited",
					ResponseBody: data.Users[userFoundId],
				})
			}
		}
	} else {
		ctx.JSON(400, data.Response{
			Success:      false,
			Messages:     "User not found !",
			ResponseBody: "",
		})
		return
	}
}
