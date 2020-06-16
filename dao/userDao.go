package dao

import (
	"chat/common"
	"chat/models"
	"chat/tools"
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
		return common.ResponseFailErr(err), userPass
	}
	if res {
		return common.QueryDataSuccess, userPass
	} else {
		return common.QueryDataFail, userPass
	}
}



//RoleQuery 角色名称查询
func (ud UserDao) RoleQuery(userName string) (string, string) {
	userRole := models.UserRole{}
	res, err := ud.Table("user").Join("INNER", "role", "user.roleId = role.roleId and user.userName = ?", userName).Get(&userRole)
	if err != nil {
		return common.ResponseFailErr(err), ""
	}
	if res {
		return common.QueryDataSuccess, userRole.RoleName
	} else {
		return common.QueryDataFail, ""
	}
}



//UserQuery user表数据查询
func (ud UserDao) UserQuery(userName string) (string, models.User) {
	user := models.User{}
	_, err := ud.Table("user").Where("userName=?", userName).Get(&user)
	if err != nil {
		return common.ResponseFailErr(err), user
	}
	return common.QueryDataSuccess, user
}
