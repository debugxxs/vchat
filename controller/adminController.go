package controller

import (
	"chat/models"
	"chat/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	service.AdminService
}
//AddUsers 添加单个用户数据
func (ac AdminController)AddUsers(c *gin.Context){
	user := models.User{}
	if err := c.ShouldBind(&user);err!=nil{
		c.JSON(200,gin.H{
			"code":200,
			"msg":fmt.Sprintln("参数解析失败",err),
		})
		return
	}
	msg,res:=ac.CheckUsers(user)
	if res{
		c.JSON(200,gin.H{
			"code":200,
			"msg":"数据插入成功",
		})
	}else {
		c.JSON(200,gin.H{
			"code":400,
			"msg":msg,
		})
	}

}
func (ac AdminController)GetUsers(c *gin.Context){

}