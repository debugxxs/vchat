package service

import (
	"chat/common"
	"chat/dao"
	"chat/models"
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


//InsertTk 插入token数据处理
func (us UserService)InsertTk(userName string,token string){
	userDao := dao.NewUserDao()
	userDao.InsertUserToken(userName,token)
}
//CheckUserOrganization 架构用户数据验证和处理
func (us UserService)CheckUserOrganization()(string,bool,interface{}){
	//1 获取组织架构下的组织列表
	//1.1 先根据组织用户来查询到相同的用户，并把他们放到同一个列表
	//2 根据组织id查询列表的值，并根据层级结构设置key
	//3 组织数据并返回
	//组织三层结构
	topData := make(map[string]interface{})
	//data := make([]interface{},0)
	userDao := dao.NewUserDao()
	msg,res,organizations :=userDao.GetAllOrganization()
	if res{
		for _,v := range organizations{
			//根据id去查询对应的用户列表
			_,_,userList :=userDao.QueryUserOrganization(v.OrganizationId)
			if v.ThreeLayer ==""{
				twoData := make(map[string]interface{})
				userNames := make([]models.UserListData,0)
				if v.TwoLayer == ""{
					for _,lv := range userList{
						users := models.UserListData{UserId: lv.UserId,UserName: lv.UserName}
						userNames = append(userNames,users)
					}
					topData[v.TOPLayer] = userNames
					//data = append(data,topData)
				}else {
					for _,lv := range userList{
						users := models.UserListData{UserId: lv.UserId,UserName: lv.UserName}
						userNames = append(userNames,users)
					}
					twoData[v.TwoLayer] = userNames
					topData[v.TOPLayer] = twoData
					//data = append(data,topData)
				}
			}else {
				userNames := make([]models.UserListData,0)
				twoData := make(map[string]interface{})
				threeData := make(map[string]interface{})
				for _,lv :=range userList{
					users := models.UserListData{UserId: lv.UserId,UserName: lv.UserName}
					userNames = append(userNames,users)
				}
				threeData[v.ThreeLayer] = userNames
				twoData[v.TwoLayer] = threeData
				topData[v.TOPLayer] = twoData
			}

		}
		return msg,res,topData
		//fmt.Println("输出结果：",data)
	}else {
		return msg,res,topData
	}
}


/*以下是密码service*/
//CheckPass 检查密码验证
func (us UserService)CheckPass(userName ,token string,userPass models.Password)(string,bool){
	userDao:=dao.NewUserDao()
	userToken := userDao.QueryUserToken(userName)
	if userToken != token{
		return common.ModifyPassErr,false
	}
	if userName == ""{
		return common.ParamsParseFail,false
	}
	userId:=userDao.QueryUserId(userName)
	return  userDao.AddUserPassWord(userId,userPass)
}
