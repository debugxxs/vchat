package service

import (
	"chat/dao"
)

type UserService struct {
}

//UserLoginFunction 用户登录处理函数
func (us UserService) UserLoginFunction(userName, passWord string) bool {
	userDao := dao.NewUserDao()
	_, userPass := userDao.UserLoginQuery(userName)
	if userPass.UserName == userName && userPass.PassWord == passWord {
		return true
	} else {
		return false
	}
}

//CheckUserRole 获取用户角色名称
func (us UserService) CheckUserRole(userName string) (string, string) {
	userDao := dao.NewUserDao()
	msg, result := userDao.RoleQuery(userName)
	if result != "" {
		return msg, result
	} else {
		return msg, ""
	}
}

//GetUserInfo 获取个人信息
func (us UserService) GetUserInfo(userName string) interface{} {
	userDao := dao.NewUserDao()
	_, user := userDao.UserQuery(userName)
	_, roleName := userDao.RoleQuery(userName)
	userData := map[string]interface{}{
		"userId":       user.UserId,
		"userName":     user.UserName,
		"userRole":     roleName,
		"userAvatar":   user.Avatar,
		"userPhone":    user.Phone,
		"userEmail":    user.Email,
		"userPosition": user.Position,
	}
	return userData
}
