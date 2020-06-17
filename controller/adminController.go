package controller

import (
	"chat/common"
	"chat/models"
	"chat/service"
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
			errMsg := common.ResponseFailErr(err)
			common.ResponseDataFail(errMsg,c)
		return
	}
	msg, res := ac.CheckUsers(user)
	if res {
		common.ResponseSuccessMsg(msg,c)
	} else {
		common.ResponseDataFail(msg,c)
	}

}
//GetUsers 查询所有用户
func (ac AdminController) GetUsers(c *gin.Context) {
	msg, QueryIsOk,userInfo := ac.CheckAllUsers()
	if QueryIsOk{
		common.ResponseSuccessData(msg,userInfo,c)
	}else {
		common.ResponseDataFail(msg,c)
	}
}

//UpDataUserInfo 根据用户id更新数据
func (ac AdminController) UpDataUserInfo(c *gin.Context) {
	//获取修改的用户id
	userId := c.Params.ByName("userId")
	user := models.User{}
	if err := c.ShouldBind(&user); err != nil {
		upDataParasErr := common.ResponseFailErr(err)
		common.ResponseDataFail(upDataParasErr,c)
	}

	//传入user层处理,检查用户是否存在
	msg, result := ac.CheckUserUpDataInfo(userId, user)
	if result {
		common.ResponseSuccessMsg(msg,c)
	} else {
		common.ResponseDataFail(msg,c)
		}
	}

//DelUserInfo 删除用户信息
func (ac AdminController) DelUserInfo(c *gin.Context) {
	userID := c.Params.ByName("userId")
	msg, result := ac.CheckDelUser(userID)
	if result {
		common.ResponseSuccessMsg(msg,c)
	} else {
		common.ResponseDataFail(msg,c)
	}
}

//GetDelUsers 获取已经被删除的用户信息
func (ac AdminController) GetDelUsers(c *gin.Context) {
	msg, queryIsOk,delUser := ac.CheckDel()
	if queryIsOk{
		common.ResponseSuccessData(msg,delUser,c)
	}else {
		common.ResponseDataFail(msg,c)
	}

}

/* 以下是role的控制器内容*/
//AddRoles role数据提交
func (ac AdminController)AddRoles(c *gin.Context){
	//获取提交的数据
	var role models.Role
	if err := c.ShouldBindJSON(&role);err!=nil{
		errParaMsg := common.ResponseFailErr(err)
		common.ResponseDataFail(errParaMsg,c)
		return
	}
	//将参数给service 层处理
	msg,result:=ac.CheckAddRole(role)
	if result {
		common.ResponseSuccessMsg(msg,c)
	}else {
		common.ResponseDataFail(msg,c)
	}
}
//QueryAllRole 查询所有roleName
func (ac AdminController)QueryAllRole(c *gin.Context){
	msg,queryIsOk,roles:=ac.CheckRoles()
	if queryIsOk{
		common.ResponseSuccessData(msg,roles,c)
	}else {
		common.ResponseDataFail(msg,c)
	}
}
//DelRoleName 删除role数据表
func (ac AdminController)DelRoleName(c *gin.Context)  {
	roleId := c.Params.ByName("roleId")
	roleNum,err:=strconv.ParseInt(roleId,10,64)
	if err !=nil{
		ParasErrMsg := common.ResponseFailErr(err)
		common.ResponseDataFail(ParasErrMsg,c)
	}
	msg,result:=ac.CheckDelRole(roleNum)
	if result{
		common.ResponseSuccessMsg(msg,c)
	}else {
		common.ResponseDataFail(msg,c)
	}

}


/*以下是Organization的控制器*/
//AddOrganization 获取客户端添加的组织架构信息参数和返回服务端数据
func (ac AdminController)AddOrganization(c *gin.Context){
	Organization := models.Organization{}
	if err := c.ShouldBind(&Organization);err!=nil{
		paraErrMsg := common.ResponseFailErr(err)
		common.ResponseDataFail(paraErrMsg,c)
	}
	//将解析内容传入service
	msg,result:=ac.CheckAddOrganization(Organization)
	if result{
		common.ResponseSuccessMsg(msg,c)
	}else {
		common.ResponseDataFail(msg,c)
	}
}
//UpDataOrganization 获取客户端修改组织架构参数和返回服务端数据
func (ac AdminController)UpDataOrganization(c *gin.Context)  {
	OrganizationId := c.Params.ByName("organizationId")
	organization :=models.Organization{}
	if err := c.ShouldBind(&organization);err !=nil{
		paramErrMsg := common.ResponseFailErr(err)
		common.ResponseDataFail(paramErrMsg,c)
		return
	}
	msg,result:=ac.CheckUpDataOrganization(OrganizationId,organization)
	if result{
		common.ResponseSuccessMsg(msg,c)
	}else {
		common.ResponseDataFail(msg,c)
	}
}
//DelOrganization 获取客户端url传入的id并返回服务器数据库数据
func (ac AdminController)DelOrganization(c *gin.Context){
	OrganizationId := c.Params.ByName("organizationId")
	msg, result :=ac.CheckDelOrganization(OrganizationId)
	if result {
		common.ResponseSuccessMsg(msg,c)
	}else {
		common.ResponseDataFail(msg,c)
	}

}
//GetAllOrganizations 返回服务端的组织架构的所有列表信息
func (ac AdminController)GetAllOrganizations(c *gin.Context){
	msg,queryIsOk,organizations:=ac.CheckAllOrganizations()
	if queryIsOk{
		common.ResponseSuccessData(msg,organizations,c)
	}else {
		common.ResponseDataFail(msg,c)
	}
}