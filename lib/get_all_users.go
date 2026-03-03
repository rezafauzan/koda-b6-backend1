package lib

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"rezafauzan/koda-b6-backend1/data"
)

func GetAllUsers(ctx *gin.Context) {
	conn, err := DatabaseHandler()
	if err != nil {
		fmt.Println(err)
	} else {
		rows, err := conn.Query(context.Background(), "SELECT users.id, user_profiles.user_avatar, user_profiles.first_name, user_profiles.last_name, user_credentials.email, user_credentials.phone, user_profiles.address, user_credentials.password, roles.role_name FROM users JOIN user_profiles ON user_profiles.user_id = users.id JOIN user_credentials ON user_credentials.user_id = users.id JOIN roles ON users.role_id = roles.id;")
		if err != nil {
			fmt.Println(err)
		} else {
			users, err := pgx.CollectRows(rows, func(user pgx.CollectableRow) (data.UserStruct, error) {
				var userData data.UserStruct
				err := user.Scan(
					&userData.Id,
					&userData.Avatar,
					&userData.Firstname,
					&userData.Lastname,
					&userData.Email,
					&userData.Phone,
					&userData.Address,
					&userData.Password,
					&userData.Role,
				)
				return userData, err
			})

			if err != nil {
				fmt.Println(err)
			} else {
				data.Users = users
				fmt.Println(data.Users)
			}

		}
	}
	ctx.JSON(200, data.Response{
		Success:      true,
		Messages:     "Success",
		ResponseBody: data.Users,
	})
}
