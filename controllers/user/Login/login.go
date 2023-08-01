package Login

import (
	"clouds.lgb24kcs.cn/controllers/user/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/user"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	req := request.LoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		user.UserApi.Login(c, req)
	}
}
