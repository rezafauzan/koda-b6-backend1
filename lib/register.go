package lib

import (
	"rezafauzan/koda-b6-backend1/data"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

func Register(ctx *gin.Context) {
	payload := data.UserStruct{}
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(400, data.Response{
			Success:      false,
			Messages:     "Failed to create users",
			ResponseBody: "",
		})
	} else {
		emailExist := 0
		for x := range data.Users {
			if data.Users[x].Email == payload.Email {
				emailExist++
			}
		}
		if emailExist == 0 {
			payload.Id = len(data.Users)
			if len(payload.Firstname) < 4 {
				ctx.JSON(400, data.Response{
					Success:      false,
					Messages:     "Firstname minimal 4 characters!",
					ResponseBody: "",
				})
				return
			}

			if len(payload.Lastname) < 4 {
				ctx.JSON(400, data.Response{
					Success:      false,
					Messages:     "Lastname minimal 4 characters!",
					ResponseBody: "",
				})
				return
			}

			if len(payload.Email) < 4 || strings.Contains(payload.Email, "@") != true {
				ctx.JSON(400, data.Response{
					Success:      false,
					Messages:     "Email minimal 4 characters and must be a valid email!",
					ResponseBody: "",
				})
				return
			}

			if len(payload.Phone) < 10 {
				ctx.JSON(400, data.Response{
					Success:      false,
					Messages:     "Phone numbers minimal 10 digits",
					ResponseBody: "",
				})
				return
			}

			if len(payload.Address) < 10 {
				ctx.JSON(400, data.Response{
					Success:      false,
					Messages:     "Address minimal 10 characters",
					ResponseBody: "",
				})
				return
			}

			if len(payload.Password) < 8 {
				ctx.JSON(400, data.Response{
					Success:      false,
					Messages:     "Password too weak minimal 8 characters!",
					ResponseBody: "",
				})
				return
			} else {
				argon := argon2.DefaultConfig()
				hash, err := argon.HashEncoded([]byte(payload.Password))
				if err != nil {
					ctx.JSON(400, data.Response{
						Success:      false,
						Messages:     "System fail to proses password!",
						ResponseBody: "",
					})
					return
				} else {
					payload.Password = string(hash)
				}
			}
			payload.Avatar = "https://i.pravatar.cc/400?img=54"
			payload.Role = "Member"
			data.Users = append(data.Users, payload)
			ctx.JSON(200, data.Response{
				Success:      true,
				Messages:     "Users created",
				ResponseBody: data.Users,
			})
		} else {
			ctx.JSON(400, data.Response{
				Success:      false,
				Messages:     "Email allready used !",
				ResponseBody: "",
			})
		}
	}
}
