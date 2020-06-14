package tools

import (
	"chat/models"
	"log"
	"xorm.io/xorm"
)
import _ "github.com/go-sql-driver/mysql"

type Orm struct {
	*xorm.Engine
}

var DbEngine *Orm

//ParsingDbConfig 解析数据库配置文件
func ParsingDbConfig(config *Config) *Orm {
	dbconfig := config.Database
	conn := dbconfig.DbUser + ":" + dbconfig.DbPass + "@/" + dbconfig.DbName + "?charset=" + dbconfig.CharSet
	engine, err := xorm.NewEngine(dbconfig.Drive, conn)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//engine.ShowSQL(dbconfig.ShowSql)
	err = engine.Sync2(new(models.User), new(models.Role), new(models.Organization), new(models.Password))
	if err != nil {
		log.Fatalf(err.Error())
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm
}
