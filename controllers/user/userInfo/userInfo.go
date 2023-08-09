package userInfo

import (
	"clouds.lgb24kcs.cn/services/user"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 用户详情
// @Success 200 {string} string json{"code","message"}
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Router  /user/v1/userinfo  [get]
// @version 1.0
func UserInfo(c *gin.Context) {
	id := utils.CtxUserId(c)
	utils.CtxUserId(c)
	user.UserInfo.UserInfo(c, id)
}
