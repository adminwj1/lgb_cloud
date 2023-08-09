package upload

import (
	"clouds.lgb24kcs.cn/controllers/user/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/user"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 用户更新
// @Success 200 {string} string json{"code","message"}
// @Param data body request.UpdataReq{} true "添加请求参数"
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Router  /user/v1/userupdata  [put]
// @version 1.0
func UserUpdata(c *gin.Context) {
	id := utils.CtxUserId(c)
	req := request.UpdataReq{}
	if err := c.ShouldBind(&req); err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		user.Updata.UserUpdata(c, req, id)
	}

}
