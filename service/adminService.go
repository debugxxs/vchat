package service

import (
	"chat/common"
	"chat/dao"
	"chat/models"
	"strconv"
)

type AdminService struct {
}
//CheckUsers 用户提交数据处理
func (as *AdminService) CheckUsers(user models.User) (string, bool) {
	adminDao := dao.NewAdminDao()
	msg, res := adminDao.AddUsers(user)
	if res != 0 {
		return msg, true
	} else {
		return msg, false
	}
}
//CheckAllUsers 处理查询所有用户信息查询
func (as AdminService) CheckAllUsers() (string, bool,[]models.UserAllData) {
	//1.先查询user表全部信息
	//2 判断对应id是否为0，如果是0，就不去查询对应数值，如果不是就返回对应的信息表
	//3处理数据信息
	adminDao := dao.NewAdminDao()
	userData := make([]models.UserAllData, 0)

	msg, isExit,users := adminDao.QueryUsers()
	if isExit{
		for _, v := range users {
			if v.UserName == "admin" {
				continue
			}
			if v.RoleId != 0 || v.OrganizationId != 0 {
				_, roleName := adminDao.QueryUserRoleName(v.UserId)
				_, Organization := adminDao.QueryUserOrganizationName(v.UserId)
				userRpData := models.UserAllData{UserId: v.UserId, UserName: v.UserName, Phone: v.Phone, Avatar: v.Avatar, Email: v.Email, Position: v.Position, RoleName: roleName, OrganizationName: Organization}
				userData = append(userData, userRpData)
			} else {
				userRpData := models.UserAllData{UserId: v.UserId, UserName: v.UserName, Phone: v.Phone, Avatar: v.Avatar, Email: v.Email, Position: v.Position, RoleName: "", OrganizationName: models.Organization{}}
				userData = append(userData, userRpData)
			}
		}
		return msg, true,userData
	}else {
		return msg,false,userData
	}

}
//CheckUserUpDataInfo 用户更新数据方法
func (as AdminService) CheckUserUpDataInfo(userid string, userInfo models.User) (string, bool) {
	adminDao := dao.NewAdminDao()
	userId, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		return common.ResponseFailErr(err), false
	}
	msg, res := adminDao.QueryUserIsExit(userId)
	if res {
		upDataMsg, res := adminDao.UpDataUserInfo(userId, userInfo)
		if res != 0 {
			return upDataMsg, true
		} else {
			return upDataMsg, false
		}
	} else {
		return msg, false
	}
}
//CheckDelUser 处理delete 方法数据
func (as AdminService) CheckDelUser(userid string) (string, bool) {
	userId, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		return common.ResponseFailErr(err), false
	}
	adminDao := dao.NewAdminDao()
	msg, result := adminDao.DelUser(userId)
	if result != 0 {
		return msg, true
	} else {
		return msg, false
	}
}
//CheckDel 删除用户信息的返回
func (as AdminService) CheckDel() (string, bool,[]models.UserAllData) {
	adminDao := dao.NewAdminDao()
	delUserData := make([]models.UserAllData, 0)
	msg,isExits, delUser := adminDao.QueryDelUser()
	if isExits{
		for _, v := range delUser {
			if v.UserName == "admin" {
				continue
			}
			if v.DeletedAt > 0 {
				_, roleName := adminDao.QueryUserRoleName(v.UserId)
				_, Organization := adminDao.QueryUserOrganizationName(v.UserId)
				userRpData := models.UserAllData{UserId: v.UserId, UserName: v.UserName, Phone: v.Phone, Avatar: v.Avatar, Email: v.Email, Position: v.Position, RoleName: roleName, DeletedAt: v.DeletedAt, OrganizationName: Organization}
				delUserData = append(delUserData, userRpData)
			}
		}
		return msg, true,delUserData
	}else {
		return msg,false,delUserData
	}

}

/*下面是role service内容*/
//CheckRole 处理用户添加数据
func (as AdminService)CheckAddRole(role models.Role)(string,bool){
	adminDao := dao.NewAdminDao()
	msg,result:=adminDao.AddRole(role)
	if result !=0{
		return msg,true
	}else {
		return msg,false
	}
}
//处理查询所有role数据
func (as AdminService)CheckRoles()(string,bool,[]models.Role){
	adminDao := dao.NewAdminDao()
	msg,queryIsOk,roles:=adminDao.QueryAllRole()
	return msg,queryIsOk,roles
}
//删除数据处理
func (as AdminService)CheckDelRole(roleId int64)(string,bool){
	adminDao := dao.NewAdminDao()
	msg,res:=adminDao.DelRoleName(roleId)
	if res !=0{
		return msg,true
	}else {
		return msg,false
	}
}


/*以下是Organization service*/
//CheckAddOrganization 处理添加组织架构信息Service
func (as AdminService)CheckAddOrganization(Organization models.Organization)(string,bool){
	adminDao := dao.NewAdminDao()
	msg,result:=adminDao.AddOrganizationData(Organization)
	if result !=0{
		return msg,true
	}else {
		return msg,false
	}
}
//CheckUpDataOrganization 处理修改组织架构service
func (as AdminService)CheckUpDataOrganization(organizationid string,Organization models.Organization)(string,bool){
	OrganizationId,err := strconv.ParseInt(organizationid,10,64)
	if err!=nil{
		return common.ResponseFailErr(err),false
	}
	adminDao:= dao.NewAdminDao()
	msg,result:=adminDao.UpDataOrganizationData(OrganizationId,Organization)
	if result !=0{
		return msg,true
	}else {
		return msg,false
	}
}
//CheckDelOrganization 处理删除组织架构service
func (as AdminService)CheckDelOrganization(organizationid string)(string,bool){
	adminDao := dao.NewAdminDao()
	OrganizationId,err := strconv.ParseInt(organizationid,10,64)
	if err !=nil{
		return common.ResponseFailErr(err),false
	}
	msg,result:=adminDao.DelOrganizationData(OrganizationId)
	if result !=0{
		return msg,true
	}else {
		return msg,false
	}
}
//CheckAllOrganizations 处理获取所有组织架构信息service
func (as AdminService)CheckAllOrganizations()(string,bool,[]models.Organization){
	adminDao :=dao.NewAdminDao()
	msg,isOk,organizations:=adminDao.GetAllOrganizationsData()
	return msg,isOk,organizations
}


