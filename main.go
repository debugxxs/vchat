package main

import (
	"chat/routers"
	"chat/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	//加载配置
	cfg, err := tools.ParsingConfig("./config/config.json")
	if err != nil {
		panic(err.Error())
	}
	//设置运行模式
	tools.ParsingDbConfig(cfg)
	app := gin.Default()
	routers.LoadRouter(app)
	_ = app.Run(cfg.AppHost + ":" + cfg.AppPort)
}
