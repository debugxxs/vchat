package controller

import (
	"chat/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service.UserService
}

func (uc UserController) IndexHandel(c *gin.Context) {
	//获取组织列表中的好友
	
}
