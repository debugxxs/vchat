package controller

import (
	"chat/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service.UserService
}

func (uc UserController)Hello(c *gin.Context)  {
	c.JSON(200,gin.H{
		"code":200,
		"data":c.FullPath(),
	})
}