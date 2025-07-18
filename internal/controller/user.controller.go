package controller

import (
	"github.com/AndyLab9596/go-ecommerce-backend-api/internal/service"
	"github.com/AndyLab9596/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserByID(c *gin.Context) {
	response.SuccessReponse(c, response.ErrCodeSuccess, []string{"tipjs", "hello"})
}
