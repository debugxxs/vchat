package routers

import (
	"chat/controller"
	"chat/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRouter(engine *gin.Engine)  {
	var (
		myjwt middleware.AuthJwt
		user controller.UserController
		adminUser controller.AdminController
	)
	allAuthMiddleware := myjwt.AuthMiddlewareFunc(myjwt.AllAuthMiddleware)
	engine.NoRoute(allAuthMiddleware.MiddlewareFunc(),myjwt.NoRouteHandler)
	engine.POST("/login",allAuthMiddleware.LoginHandler)
	userApi:= engine.Group("/user")
	{
		userApi.GET("/refresh_token",allAuthMiddleware.RefreshHandler)
	}
	userApi.Use(allAuthMiddleware.MiddlewareFunc())
	{
		userApi.GET("/hello",user.Hello)
	}
	adminAuthMinddleware := myjwt.AuthMiddlewareFunc(myjwt.AdminAuthMiddleware)
	adminApi := engine.Group("/api/v1")
	adminApi.Use(adminAuthMinddleware.MiddlewareFunc())
	{
		adminApi.GET("/users",adminUser.AddUsers)
	}

}
