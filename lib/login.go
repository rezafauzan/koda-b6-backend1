package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"rezafauzan/koda-b6-backend1/data"
)

func Login(ctx *gin.Context) {
	payload := data.UserStruct{}
	err := ctx.ShouldBindJSON(&payload)
	if len(payload.Email) < 1 {
		if err != nil {
			ctx.JSON(400, data.Response{
				Success:      false,
				Messages:     "Login failed",
				ResponseBody: "",
			})
			return
		} else {
			for x := range data.Users {
				if data.Users[x].Email == payload.Email {
					correct, err := argon2.VerifyEncoded([]byte(payload.Password), []byte(data.Users[x].Password))
					if err != nil {
						ctx.JSON(400, data.Response{
							Success:      false,
							Messages:     "Email or password wrong!",
							ResponseBody: "",
						})
						return
					} else if correct {
						data.User = data.Users[x]
						data.User.Password = "Hidden"
						ctx.JSON(200, data.Response{
							Success:      false,
							Messages:     "Login success! wellcome back " + data.Users[x].Firstname + " " + data.Users[x].Lastname,
							ResponseBody: data.User,
						})
						return
					}
				} else {
					ctx.JSON(400, data.Response{
						Success:      false,
						Messages:     "Email or password wrong!",
						ResponseBody: "",
					})
					return
				}
			}
		}
	} else {
		ctx.JSON(200, data.Response{
			Success:      false,
			Messages:     "You allready logged in!!",
			ResponseBody: data.User,
		})
		return
	}
}
