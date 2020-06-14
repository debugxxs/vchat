package dao

import (
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
		return "数据查询失败", 0
	}
	if qureyres {
		return "用户已经存在", 0
	} else {
		res, err := ad.Table("user").InsertOne(&userInfo)
		if err != nil {
			return fmt.Sprintln("数据查询失败", err), 0
		}
		return "数据查询成功", res
	}
}
//查询所有用户信息
func (ad AdminDao) QueryUsers() (string, []models.User) {
	users := make([]models.User, 0)
	err := ad.Table("user").Distinct("userId", "userName", "phone", "eMail", "avatar", "position", "roleId", "organizationId").Find(&users)
	if err != nil {
		return fmt.Sprintln("数据查询失败", err), users
	}
	return "数据查询成功", users
}
//QueryUserRoleName 查询用户角色
func (ad AdminDao) QueryUserRoleName(userId int64) (string, string) {
	userRole := models.UserRole{}
	res, err := ad.Table("user").Join("INNER", "role", "user.roleId = role.roleId and user.userId = ?", userId).Get(&userRole)
	if err != nil {
		return fmt.Sprintln("数据查询失败", err), userRole.RoleName
	}
	if res {
		return "数据查询成功", userRole.RoleName
	} else {
		return "用户角色不存在", userRole.RoleName
	}
}
//QueryUserOrganizationName 查询用户组织架构信息
func (ad AdminDao) QueryUserOrganizationName(userId int64) (string, models.Organization) {
	userOrganization := models.UserOrganization{}
	res, err := ad.Table("user").Join("INNER", "organization", "Organization.OrganizationId = user.OrganizationId and user.userId = ?", userId).Get(&userOrganization)
	if err != nil {
		return fmt.Sprintln("数据查询失败", err), models.Organization{}
	}
	Organization := models.Organization{TOPLayer: userOrganization.TOPLayer, TwoLayer: userOrganization.TwoLayer, ThreeLayer: userOrganization.ThreeLayer}
	if res {
		return "数据查询成功", Organization
	} else {
		return "组织架构不存在", Organization
	}
}
//QueryUserIsExit 查询用户信息是否存在
func (ad AdminDao) QueryUserIsExit(userId int64) (string, bool) {
	user := models.User{}
	res, err := ad.Table("user").Where("userId = ?", userId).Get(&user)
	if err != nil {
		return "查询数据出错", false
	}
	if res {
		return "数据查询成功", true
	} else {
		return "用户不存在", false
	}
}
//UpDataUserInfo 更新用户信息
func (ad AdminDao) UpDataUserInfo(userId int64, userInfo models.User) (string, int64) {
	res, err := ad.Table("user").ID(userId).Update(&userInfo)
	if err != nil {
		return fmt.Sprintln("数据更新失败", err), res
	} else {
		return "数据更新成功", res
	}
}
//DelUser 删除用户
func (ad AdminDao) DelUser(userId int64) (string, int64) {
	user := models.User{UserId: userId}
	res, err := ad.Table("user").ID(userId).Delete(&user)
	if err != nil {
		return fmt.Sprintln("数据删除失败", err), res
	}
	if res != 0 {
		return "数据删除成功", res
	} else {
		return "数据删除失败", 0
	}
}
//QueryDelUser 查询删除用户
func (ad AdminDao) QueryDelUser() (string, []models.User) {
	users := make([]models.User, 0)
	err := ad.Table("user").Distinct("userId", "userName", "phone", "eMail", "avatar", "position", "roleId", "organizationId", "deleted_at").Unscoped().Find(&users)
	fmt.Println(users)
	if err != nil {
		return fmt.Sprintln("数据查询失败", err), users
	}
	return "数据查询成功", users
}

/*以下是role表数据库操作*/
//AddRole 用户表表添加数据
func (ad AdminDao)AddRole(role models.Role)(string,int64){
	res,err:=ad.Table("role").Insert(&role)
	if err !=nil{
		return fmt.Sprintln("数据插入失败",err),res
	}
	if res !=0{
		return "数据添加成功",res
	}else {
		return "数据添加失败",res
	}
}
//QueryAllRole 查询全部role表数据
func (ad AdminDao)QueryAllRole()(string,[]models.Role){
	roles := make([]models.Role,0)
	 err := ad.Table("role").Distinct("roleName","roleId").Find(&roles)
	 if err!=nil{
	 	return "用户查询失败",roles
	 }
	 return "数据查询成功",roles
}
//DelRoleName 删除role用户
func (ad AdminDao)DelRoleName(roleId int64)(string,int64){
	role := models.Role{RoleId: roleId}
	fmt.Println(roleId)
	result,err:=ad.Table("role").Where("roleId = ?",roleId).Get(&role)
	if err !=nil{
		return fmt.Sprintln("提交参数有误",err),0
	}
	if result{
		res,err:=ad.Table("role").Delete(&role)
		if err !=nil{
			return fmt.Sprintln("删除数据失败",err),res
		}
		if res !=0{
			return "数据删除成功",res
		}else {
			return "数据删除失败",res
		}
	}else {
		return "Id不存在",0
	}


}

/*以下是OrganizationData 表数据库操作*/
//AddOrganizationData 添加组织架构信息
func (ad AdminDao)AddOrganizationData(Organization models.Organization)(string,int64){
	res,err:=ad.Table("organization").Insert(Organization)
	if err !=nil{
		return fmt.Sprintln("参数错误",err),res
	}
	if res !=0{
		return "数据添加成功",res
	}else {
		return "数据添加失败",res
	}
}
//UpDataOrganizationData 修改组织架构信息
func (ad AdminDao)UpDataOrganizationData(organizationId int64,organization models.Organization)(string,int64){
	res,err:=ad.ID(organizationId).Update(&organization)
	if err !=nil{
		return fmt.Sprintln("更新数据出错",err),res
	}
	if res !=0{
		return "数据更新成功",res
	}else {
		return "数据更新失败",res
	}
}
//DelOrganizationData 删除组织架构表格
func (ad AdminDao)DelOrganizationData(organizationId int64)(string,int64) {
	organization := models.Organization{OrganizationId: organizationId}
	res,err:=ad.Table("organization").Delete(&organization)
	if err !=nil{
		return fmt.Sprintln("删除数据出错",err),res
	}
	if res !=0{
		return "数据删除成功",res
	}else {
		return "数据删除失败",res
	}
}
//GetAllOrganizationsData 获取组织架构数据
func (ad AdminDao)GetAllOrganizationsData()(string,[]models.Organization)  {
	organizations := make([]models.Organization,0)
	err:=ad.Table("organization").Distinct("organizationId","topLayer","twoLayer","threeLayer").Find(&organizations)
	if err !=nil{
		return fmt.Sprintln("查询数据失败",err),organizations
	}
	return "数据查询成功",organizations

}

