package middleware

import (
	"chat/models"
	"chat/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"net/http"
	"time"
)

type AuthJwt struct {
	service.UserService
}
type AuthRoleFunction func(data interface{}, c *gin.Context) bool

var (
	identityKey    string = "UserName"
	GlobalUserName string
)

func (aj *AuthJwt) AuthMiddlewareFunc(authRoleFunc AuthRoleFunction) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Minute * 30,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(context *gin.Context) interface{} {
			claims := jwt.ExtractClaims(context)
			return &models.User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//绑定参数
			userPass := models.UserPass{}
			if err := c.ShouldBind(&userPass); err != nil {
				logger.Fatalln("参数绑定错误", err)
			}
			GlobalUserName = userPass.UserName
			if aj.UserLoginFunction(userPass.UserName, userPass.PassWord) {
				return &models.User{
					UserName: userPass.UserName,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: authRoleFunc,
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			userData := aj.GetUserInfo(GlobalUserName)
			c.JSON(http.StatusOK, gin.H{
				"code":  code,
				"msg":   "登录成功",
				"token": token,
				"timer": expire.Format(time.RFC3339),
				"data":  userData,
			})
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		logger.Fatalln("初始化错误", err)
	}
	return authMiddleware
}
func (aj AuthJwt) AllAuthMiddleware(data interface{}, c *gin.Context) bool {
	return true
}
func (aj AuthJwt) NoRouteHandler(c *gin.Context) {
	c.JSON(404, gin.H{"code": 404, "message": "访问的路径不存在"})
}
func (aj AuthJwt) AdminAuthMiddleware(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.User); ok {
		_, roleName := aj.CheckUserRole(v.UserName)
		switch roleName {
		case "系统管理员":
			return true
		case "普通用户":
			return false
		default:
			return false
		}
	}

	return false
}
