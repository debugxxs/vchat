package controller

import (
	"chat/common"
	"chat/models"
	"chat/service"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserController struct {
	service.UserService
}

func (uc UserController)IndexHandel(c *gin.Context) {
	//获取组织列表中的好友
	msg,res,userOrganization:=uc.CheckUserOrganization()
	if res{
		common.ResponseSuccessData(msg,userOrganization,c)
	}else {
		common.ResponseDataFail(msg,c)
	}

	
}

/*密码控制器*/
//AddPass 第一次登录设置密码控制器和修改密码
func (uc UserController)AddAndModifyPass(c *gin.Context){
	//获取前端传入的数据
	var userPass models.UserPass
	if err := c.ShouldBindJSON(&userPass);err!=nil{
		errPassMsg := common.ResponseFailErr(err)
		common.ResponseDataFail(errPassMsg,c)
	}
	auth := c.Request.Header.Get("Authorization")
	handelList := strings.Split(auth," ")
	tokenStr := handelList[1]
	userName := userPass.UserName
	msg,res:=uc.CheckPass(userName,tokenStr,userPass.Password)
	if res{
		common.ResponseSuccessMsg(msg,c)
	}else {
		common.ResponseDataFail(msg,c)
	}

}
