package upload

import (
	"clouds.lgb24kcs.cn/controllers/user/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/user"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

func UserUpdata(c *gin.Context) {
	id := utils.CtxUserId(c)
	req := request.UpdataReq{}
	if err := c.ShouldBind(&req); err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		user.Updata.UserUpdata(c, req, id)
	}

}
