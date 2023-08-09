package Login

import (
	"clouds.lgb24kcs.cn/controllers/user/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/user"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

// @Summary  用户登录
// @Success 200 {string} string json{"code","message"}
// @Param data body request.LoginReq{} true "添加请求参数"
// @Router  /user/v1/login  [post]
// @version 1.0
func Login(c *gin.Context) {
	req := request.LoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		user.UserApi.Login(c, req)
	}
}
