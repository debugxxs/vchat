package controller

import (
	"chat/common"
	"chat/models"
	"chat/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
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
	fmt.Println(auth)
	taken := strings.Split(auth," ")
	for _,v := range taken{
		fmt.Println(v)
	}
	userName := userPass.UserName
	msg,res:=uc.CheckPass(userName,userPass.Password)
	if res{
		common.ResponseSuccessMsg(msg,c)
	}else {
		common.ResponseDataFail(msg,c)
	}

}
