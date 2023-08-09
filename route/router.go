package route

import (
	"clouds.lgb24kcs.cn/controllers/bucket/add"
	"clouds.lgb24kcs.cn/controllers/bucket/del"
	"clouds.lgb24kcs.cn/controllers/bucket/detail"
	"clouds.lgb24kcs.cn/controllers/bucket/list"
	"clouds.lgb24kcs.cn/controllers/catalogue/create"
	"clouds.lgb24kcs.cn/controllers/catalogue/detailsobject"
	"clouds.lgb24kcs.cn/controllers/user/Login"
	"clouds.lgb24kcs.cn/controllers/user/Register"
	"clouds.lgb24kcs.cn/controllers/user/upload"
	"clouds.lgb24kcs.cn/controllers/user/userInfo"
	_ "clouds.lgb24kcs.cn/docs"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Cors())
	user := engine.Group("user/v1")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	{
		user.POST("login", Login.Login)
		user.POST("register", Register.Register)

		user.Use(middleware.ChecKToken())
		{
			user.GET("userinfo", userInfo.UserInfo)
			user.PUT("userupdata", upload.UserUpdata)

		}
	}
	bucket := engine.Group("bucket/v1")
	bucket.Use(middleware.ChecKToken())
	{
		bucket.POST("createbucket", add.Add)
		bucket.GET("list", list.List)
		bucket.GET("detail", detail.DetailBucket)
		bucket.DELETE("del", del.DELBucket)
	}
	catalogue := engine.Group("catalogue/v1")
	catalogue.Use(middleware.ChecKToken())
	{
		catalogue.POST("create", create.Create)
		catalogue.GET("detail", detailsobject.DetailObject)
	}

	return engine
}

func Start() {
	router := StartRouter()
	router.Run(":" + global.APP.Configuration.Server.Port)
}
