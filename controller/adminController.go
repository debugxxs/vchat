package controller

import (
	"chat/models"
	"chat/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AdminController struct {
	service.AdminService
}

//AddUsers 添加单个用户数据
func (ac AdminController) AddUsers(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  fmt.Sprintln("参数解析失败", err),
		})
		return
	}
	msg, res := ac.CheckUsers(user)
	if res {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "数据插入成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  msg,
		})
	}

}
//GetUsers 查询所有用户
func (ac AdminController) GetUsers(c *gin.Context) {
	msg, userInfo := ac.CheckAllUsers()
	switch msg {
	case "数据查询成功":
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  msg,
			"data": userInfo,
		})
	case "数据查询失败":
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  msg,
		})
	default:
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  msg,
		})
	}
}
//UpDataUserInfo 根据用户id更新数据
func (ac AdminController) UpDataUserInfo(c *gin.Context) {
	//获取修改的用户id
	userId := c.Params.ByName("userId")
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数解析失败",
		})
	}

	//传入user层处理,检查用户是否存在
	msg, result := ac.CheckUserUpDataInfo(userId, user)
	if result {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  msg,
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  msg,
		})
	}

}
//DelUserInfo 删除用户信息
func (ac AdminController) DelUserInfo(c *gin.Context) {
	userID := c.Params.ByName("userId")
	msg, result := ac.CheckDelUser(userID)
	if result {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  msg,
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  msg,
		})
	}
}
//GetDelUsers 获取已经被删除的用户信息
func (ac AdminController) GetDelUsers(c *gin.Context) {
	msg, delUser := ac.CheckDel()
	switch msg {
	case "数据查询成功":
		c.JSON(200, gin.H{"code": 200, "msg": msg, "data": delUser})
	case "数据查询失败":
		c.JSON(400, gin.H{"code": 400, "msg": msg})
	}

}

/* 以下是role的控制器内容*/
//AddRoles role数据提交
func (ac AdminController)AddRoles(c *gin.Context){
	//获取提交的数据
	role := models.Role{}
	if err := c.ShouldBind(&role);err !=nil{
		c.JSON(400,gin.H{
			"code":400,
			"msg":fmt.Sprintln("参数解析失败",err),
		})
	}
	//将参数给service 层处理
	msg,result:=ac.CheckAddRole(role)
	if result {
		c.JSON(200,gin.H{"code":200,"msg":msg})
	}else {
		c.JSON(400,gin.H{"code":400,"msg":msg})
	}
}
//QueryAllRole 查询所有roleName
func (ac AdminController)QueryAllRole(c *gin.Context){
	msg,roles:=ac.CheckRoles()
	switch msg {
	case "数据查询成功":
		c.JSON(200,gin.H{"code":200,"msg":msg,"data":roles})
	case "数据查询失败":
		c.JSON(400,gin.H{"code":400,"msg":msg})
	default:
		c.JSON(400,gin.H{"code":400,"msg":msg})

	}
}
//DelRoleName 删除role数据表
func (ac AdminController)DelRoleName(c *gin.Context)  {
	roleId := c.Params.ByName("roleId")
	roleNum,err:=strconv.ParseInt(roleId,10,64)
	if err !=nil{
		c.JSON(400,gin.H{"code":400,"msg":fmt.Sprintln("数据解析失败",err)})
	}
	msg,result:=ac.CheckDelRole(roleNum)
	if result{
		c.JSON(200,gin.H{"code":200,"msg":msg})
	}else {
		c.JSON(400,gin.H{"code":400,"msg":msg})
	}

}


/*以下是Organization的控制器*/
//AddOrganization 获取客户端添加的组织架构信息参数和返回服务端数据
func (ac AdminController)AddOrganization(c *gin.Context){
	Organization := models.Organization{}
	if err := c.ShouldBind(&Organization);err!=nil{
		c.JSON(400,gin.H{"code":400,"msg":fmt.Sprintln("参数解析失败",err)})
	}
	//将解析内容传入service
	msg,result:=ac.CheckAddOrganization(Organization)
	if result{
		c.JSON(200,gin.H{"code":200,"msg":msg})
	}else {
		c.JSON(400,gin.H{"code":400,"msg":msg})
	}
}
//UpDataOrganization 获取客户端修改组织架构参数和返回服务端数据
func (ac AdminController)UpDataOrganization(c *gin.Context)  {
	OrganizationId := c.Params.ByName("organizationId")
	organization :=models.Organization{}
	if err := c.ShouldBind(&organization);err !=nil{
		c.JSON(400,gin.H{"code":400,"msg":fmt.Sprintln("参数解析错误",err)})
		return
	}
	msg,result:=ac.CheckUpDataOrganization(OrganizationId,organization)
	if result{
		c.JSON(200,gin.H{"code":200,"msg":msg})
	}else {
		c.JSON(400,gin.H{"code":400,"msg":msg})
	}
}
//DelOrganization 获取客户端url传入的id并返回服务器数据库数据
func (ac AdminController)DelOrganization(c *gin.Context){
	OrganizationId := c.Params.ByName("organizationId")
	msg, result :=ac.CheckDelOrganization(OrganizationId)
	if result {
		c.JSON(200,gin.H{"code":200,"msg":msg})
	}else {
		c.JSON(400,gin.H{"code":400,"msg":msg})
	}

}
//GetAllOrganizations 返回服务端的组织架构的所有列表信息
func (ac AdminController)GetAllOrganizations(c *gin.Context){
	msg,organizations:=ac.CheckAllOrganizations()
	switch msg {
	case "数据查询成功":
		c.JSON(200,gin.H{"code":200,"msg":msg,"data":organizations})
	default:
		c.JSON(400,gin.H{"code":400,"msg":msg})

	}
}
