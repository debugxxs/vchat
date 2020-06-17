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



func(ud UserDao)QueryUserId(userName string)int64{
	var user models.User
	ud.Table("user").Where("userName = ? ",userName).Get(&user)
	return user.UserId
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

/*以下是密码表数据库操作*/

func (ud UserDao)AddUserPassWord(userId int64,userPass models.Password)(string,bool){
	//先查询密码表是否存在相同密码
	var userPw models.Password
	res,err:=ud.Table("password").Where("passWord = ?",userPass.PassWord).Get(&userPw)
	if err !=nil{
		errMsg := common.ResponseFailErr(err)
		return errMsg,false
	}
	if res{
		//如果存在就给user表的passID赋值
		user := models.User{PassId: userPw.PassId}
		res,err:=ud.Table("user").ID(userId).Update(&user)
		if err !=nil{
			insertErrMsg := common.ResponseFailErr(err)
			return insertErrMsg,false
		}
		if res!=0{
			return common.InsertDataSuccess,true
		}else {
			return common.InsertDataFail,false
		}
	}else {
		res,err :=ud.Table("password").Insert(&userPass)
		if err != nil{
			errMsg := common.ResponseFailErr(err)
			return errMsg,false
		}
		if res !=0{
			password := userPass.PassWord
			passId:=ud.QueryUserPass(password)
			user := models.User{PassId: passId}
			_,_=ud.Table("user").ID(userId).Update(&user)
			return common.InsertDataSuccess,true
		}else {
			return common.InsertDataFail,false
		}
		//还需要分配id给用户，还是需要查询一次密码
	}
}

func (ud UserDao)QueryUserPass(password string)int64{
	var userPass models.Password
	ud.Table("password").Where("passWord = ? ",password).Get(&userPass)
	return userPass.PassId
}
