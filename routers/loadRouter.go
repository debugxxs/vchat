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
		userApi.GET("/index",user.IndexHandel)
	}
	adminAuthMiddleware := myjwt.AuthMiddlewareFunc(myjwt.AdminAuthMiddleware)
	adminApi := engine.Group("/api/v1")
	adminApi.Use(adminAuthMiddleware.MiddlewareFunc())
	{	//用户表的增删改查
		adminApi.POST("/users",adminUser.AddUsers)
		adminApi.GET("/users",adminUser.GetUsers)
		adminApi.PUT("/users/:userId",adminUser.UpDataUserInfo)
		adminApi.DELETE("/user/:userId",adminUser.DelUserInfo)
		adminApi.GET("/users/delUser",adminUser.GetDelUsers)
		//role 表增删改查
		adminApi.POST("/roles",adminUser.AddRoles)
		adminApi.GET("/roles",adminUser.QueryAllRole)
		adminApi.DELETE("/roles/:roleId",adminUser.DelRoleName)
		//组织架构路由 增删改查
		adminApi.POST("/organizations",adminUser.AddOrganization)
		adminApi.PUT("/organization/:organizationId",adminUser.UpDataOrganization)
		adminApi.DELETE("/organization/:organizationId",adminUser.DelOrganization)
		adminApi.GET("/organizations",adminUser.GetAllOrganizations)
	}

}
