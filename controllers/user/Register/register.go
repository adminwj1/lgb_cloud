package Register

import (
	"clouds.lgb24kcs.cn/controllers/user/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/user"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	req := request.RegisterReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		user.RegisterAPI.Register(c, req)
	}
}
