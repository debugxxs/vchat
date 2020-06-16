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
	//data:["top":{"two":{"three":users}}]
	//map[str]map[str]map[str]string
	//var data map[string]map[string]map[string]string
	//data := make(map[string]map[string]map[string]string,0)
	
}
