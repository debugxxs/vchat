package models

type User struct {
	UserId int64 `json:"userId" xorm:"pk autoincr userId"`
	UserName	string	`json:"userName"xorm:"varchar(16) userName"`
	Phone	string	`json:"phone"xorm:"varchar(12) phone"`
	Email 	string	`json:"eMail" xorm:"varchar(64) eMail"`
	Avatar 	string	`json:"avatar" xorm:"varchar(64) avatar"`
	LoginTime string	`json:"loginTime"xorm:"updated"`
	PassId	int64		`json:"passId" xorm:"passId"`
	RoleId	int64	`json:"roleId" xorm:"roleId"`
	OrganizationId int64	`json:"organizationid"xorm:"organizationId"`
}
 type Password struct {
 	PassId int64	`json:"pass_id" xorm:"pk autoincr passId"`
 	PassWord string	`json:"passWord"xorm:"varchar(64) passWord"`
 }
 type UserPass struct {
 	User	`xorm:"extends"`
 	Password	`xorm:"extends"`
 }