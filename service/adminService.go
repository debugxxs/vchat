package service

import (
	"chat/dao"
	"chat/models"
)

type AdminService struct {

}
//CheckUsers 用户提交数据处理
func (as *AdminService)CheckUsers(user models.User)(string,bool){
	adminDao := dao.NewAdminDao()
	msg,res:=adminDao.AddUsers(user)
	if res !=0{
		return msg,true
	}else {
		return msg,false
	}
}
func (as AdminService)CheckAllUsers()interface{}{
//参数返回data:={{},{},{}}
}