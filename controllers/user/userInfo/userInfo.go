package userInfo

import (
	"clouds.lgb24kcs.cn/services/user"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	id := utils.CtxUserId(c)
	utils.CtxUserId(c)
	user.UserInfo.UserInfo(c, id)
}
