package handlers

import (
	"net/http"
	"rezafauzan/koda-b6-backend1/internal/dto"
	"rezafauzan/koda-b6-backend1/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler{
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) GetAll(ctx *gin.Context){
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Messages: "List all users",
		Results: handler.service.GetAll(),
	})
}
