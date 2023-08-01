package route

import (
	"clouds.lgb24kcs.cn/controllers/user/Login"
	"clouds.lgb24kcs.cn/controllers/user/Register"
	"clouds.lgb24kcs.cn/controllers/user/upload"
	"clouds.lgb24kcs.cn/controllers/user/userInfo"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/middleware"
	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	engine := gin.Default()
	user := engine.Group("user/v1")
	{
		user.POST("login", Login.Login)
		user.POST("register", Register.Register)

		user.Use(middleware.ChecKToken())
		{
			user.GET("userinfo", userInfo.UserInfo)
			user.PUT("userupdata", upload.UserUpdata)

		}
	}
	return engine
}

func Start() {
	router := StartRouter()
	router.Run(":" + global.APP.Configuration.Server.Port)
}
