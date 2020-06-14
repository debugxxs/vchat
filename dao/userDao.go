package dao

import (
	"chat/models"
	"chat/tools"
	"fmt"
)

type UserDao struct {
	*tools.Orm
}

func NewUserDao() *UserDao {
	return &UserDao{tools.DbEngine}
}

//UserLoginQuery 用户密码登录查询
func (ud UserDao) UserLoginQuery(userName string) (string, models.UserPass) {
	userPass := models.UserPass{}
	res, err := ud.Table("user").Join("INNER", "password", "user.passId = password.passId and user.userName = ?", userName).Get(&userPass)
	if err != nil {
		return fmt.Sprintf("查询失败", err), userPass
	}
	if res {
		return "查询成功", userPass
	} else {
		return "查询失败", userPass
	}
}

//RoleQuery 角色名称查询
func (ud UserDao) RoleQuery(userName string) (string, string) {
	userRole := models.UserRole{}
	res, err := ud.Table("user").Join("INNER", "role", "user.roleId = role.roleId and user.userName = ?", userName).Get(&userRole)
	if err != nil {
		return fmt.Sprintln("查询数据出错", err), ""
	}
	if res {
		return "数据查询成功", userRole.RoleName
	} else {
		return "数据查询失败", ""
	}
}

//UserQuery user表数据查询
func (ud UserDao) UserQuery(userName string) (string, models.User) {
	user := models.User{}
	_, err := ud.Table("user").Where("userName=?", userName).Get(&user)
	if err != nil {
		return fmt.Sprintln("数据查询出错", err), user
	}
	return "数据查询成功", user
}
