package dao

import (
	"chat/common"
	"chat/models"
	"chat/tools"
	"fmt"
)

type AdminDao struct {
	*tools.Orm
}

func NewAdminDao() *AdminDao {
	return &AdminDao{tools.DbEngine}
}

//添加用户信息
func (ad *AdminDao) AddUsers(userInfo models.User) (string, int64) {
	user := models.User{}
	qureyres, err := ad.Table("user").Where("userName = ?", userInfo.UserName).Get(&user)
	if err != nil {
		return common.ResponseFailErr(err), 0
	}
	if qureyres {
		return common.UserIsExits, 0
	} else {
		res, err := ad.Table("user").InsertOne(&userInfo)
		if err != nil {
			return common.ResponseFailErr(err), 0
		}
		if res !=0{
			return common.InsertDataSuccess, res
		}else {
			return common.InsertDataFail,res
		}

	}
}
//查询所有用户信息
func (ad AdminDao) QueryUsers() (string, bool,[]models.User) {
	users := make([]models.User, 0)
	err := ad.Table("user").Distinct("userId", "userName", "phone", "eMail", "avatar", "position", "roleId", "organizationId").Find(&users)
	if err != nil {
		return common.ResponseFailErr(err),false, users
	}
	return common.QueryDataSuccess,true, users
}


//QueryUserRoleName 查询用户角色
func (ad AdminDao) QueryUserRoleName(userId int64) (string, string) {
	userRole := models.UserRole{}
	res, err := ad.Table("user").Join("INNER", "role", "user.roleId = role.roleId and user.userId = ?", userId).Get(&userRole)
	if err != nil {
		return common.ResponseFailErr(err), userRole.RoleName
	}
	if res {
		return common.QueryDataSuccess, userRole.RoleName
	} else {
		return common.RoleNameNotExits, userRole.RoleName
	}
}


//QueryUserOrganizationName 查询用户组织架构信息
func (ad AdminDao) QueryUserOrganizationName(userId int64) (string, models.Organization) {
	userOrganization := models.UserOrganization{}
	res, err := ad.Table("user").Join("INNER", "organization", "Organization.OrganizationId = user.OrganizationId and user.userId = ?", userId).Get(&userOrganization)
	if err != nil {
		return common.ResponseFailErr(err), models.Organization{}
	}
	Organization := models.Organization{TOPLayer: userOrganization.TOPLayer, TwoLayer: userOrganization.TwoLayer, ThreeLayer: userOrganization.ThreeLayer}
	if res {
		return common.QueryDataSuccess, Organization
	} else {
		return common.OrganizationNotExits, Organization
	}
}


//QueryUserIsExit 查询用户信息是否存在
func (ad AdminDao) QueryUserIsExit(userId int64) (string, bool) {
	user := models.User{}
	res, err := ad.Table("user").Where("userId = ?", userId).Get(&user)
	if err != nil {
		return common.ResponseFailErr(err), false
	}
	if res {
		return common.QueryDataSuccess, true
	} else {
		return common.UserIsNotExits, false
	}
}


//UpDataUserInfo 更新用户信息
func (ad AdminDao) UpDataUserInfo(userId int64, userInfo models.User) (string, int64) {
	res, err := ad.Table("user").ID(userId).Update(&userInfo)
	if err != nil {
		return common.ResponseFailErr(err), res
	} else {
		return common.UpDataSuccess, res
	}
}


//DelUser 删除用户
func (ad AdminDao) DelUser(userId int64) (string, int64) {
	user := models.User{UserId: userId}
	res, err := ad.Table("user").ID(userId).Delete(&user)
	if err != nil {
		return common.ResponseFailErr(err), res
	}
	if res != 0 {
		return common.DelDataSuccess, res
	} else {
		return common.DelDataFail, 0
	}
}


//QueryDelUser 查询删除用户
func (ad AdminDao) QueryDelUser() (string, bool,[]models.User) {
	users := make([]models.User, 0)
	err := ad.Table("user").Distinct("userId", "userName", "phone", "eMail", "avatar", "position", "roleId", "organizationId", "deleted_at").Unscoped().Find(&users)
	if err != nil {
		return common.ResponseFailErr(err),false, users
	}
	return common.QueryDataSuccess,true, users
}



/*以下是role表数据库操作*/
//AddRole 用户表表添加数据
func (ad AdminDao)AddRole(role models.Role)(string,int64){
	roles := models.Role{}
	isExits,err :=ad.Table("role").Where("roleName = ?",role.RoleName).Get(&roles)
	if err !=nil{
		return common.ResponseFailErr(err),0
	}
	if isExits{
		return common.RoleNameExits,0
	}else {
		res,err:=ad.Table("role").Insert(&role)
		if err !=nil{
			return common.ResponseFailErr(err),res
		}
		if res !=0{
			return common.InsertDataSuccess,res
		}else {
			return common.InsertDataFail,res
		}
	}

}


//QueryAllRole 查询全部role表数据
func (ad AdminDao)QueryAllRole()(string,bool,[]models.Role){
	roles := make([]models.Role,0)
	 err := ad.Table("role").Distinct("roleName","roleId").Find(&roles)
	 if err!=nil{
	 	return common.QueryDataFail,false,roles
	 }
	 return common.QueryDataSuccess,true,roles
}


//DelRoleName 删除role用户
func (ad AdminDao)DelRoleName(roleId int64)(string,int64){
	role := models.Role{RoleId: roleId}
	fmt.Println(roleId)
	result,err:=ad.Table("role").Where("roleId = ?",roleId).Get(&role)
	if err !=nil{
		return common.ResponseFailErr(err),0
	}
	if result{
		res,err:=ad.Table("role").Delete(&role)
		if err !=nil{
			return common.ResponseFailErr(err),res
		}
		if res !=0{
			return common.DelDataSuccess,res
		}else {
			return common.DelDataFail,res
		}
	}else {
		return common.IdNotExits,0
	}


}

/*以下是OrganizationData 表数据库操作*/
//AddOrganizationData 添加组织架构信息
func (ad AdminDao)AddOrganizationData(Organization models.Organization)(string,int64){
	res,err:=ad.Table("organization").Insert(Organization)
	if err !=nil{
		return common.ResponseFailErr(err),res
	}
	if res !=0{
		return common.InsertDataSuccess,res
	}else {
		return common.InsertDataFail,res
	}
}


//UpDataOrganizationData 修改组织架构信息
func (ad AdminDao)UpDataOrganizationData(organizationId int64,organization models.Organization)(string,int64){
	res,err:=ad.ID(organizationId).Update(&organization)
	if err !=nil{
		return common.ResponseFailErr(err),res
	}
	if res !=0{
		return common.UpDataSuccess,res
	}else {
		return common.UpDataFail,res
	}
}


//DelOrganizationData 删除组织架构表格
func (ad AdminDao)DelOrganizationData(organizationId int64)(string,int64) {
	organization := models.Organization{OrganizationId: organizationId}
	res,err:=ad.Table("organization").Delete(&organization)
	if err !=nil{
		return common.ResponseFailErr(err),res
	}
	if res !=0{
		return common.DelDataSuccess,res
	}else {
		return common.DelDataFail,res
	}
}


//GetAllOrganizationsData 获取组织架构数据
func (ad AdminDao)GetAllOrganizationsData()(string,bool,[]models.Organization)  {
	organizations := make([]models.Organization,0)
	err:=ad.Table("organization").Distinct("organizationId","topLayer","twoLayer","threeLayer").Find(&organizations)
	if err !=nil{
		return common.ResponseFailErr(err),false,organizations
	}
	return common.QueryDataSuccess,true,organizations

}


