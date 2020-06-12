package dao

import (
	"chat/models"
	"chat/tools"
	"fmt"
)

type AdminDao struct {
	*tools.Orm
}

func NewAdminDao()*AdminDao{
	return &AdminDao{tools.DbEngine}
}
//添加用户信息
func (ad *AdminDao)AddUsers(userInfo models.User)(string,int64){
	user := models.User{}
	qureyres,err:=ad.Table("user").Where("userName = ?",userInfo.UserName).Get(&user)
	if err !=nil{
		return "数据查询失败",0
	}
	if qureyres {
		return "用户已经存在",0
	}else {
		res,err:=ad.Table("user").InsertOne(&userInfo)
		if err !=nil{
			return fmt.Sprintln("数据查询失败",err),0
		}
		return "数据查询成功",res
	}
}