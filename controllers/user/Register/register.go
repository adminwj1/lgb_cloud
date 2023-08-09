package Register

import (
	"clouds.lgb24kcs.cn/controllers/user/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/user"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

// @Summary  用户注册接口
// @Success 200 {string} string json{"code","message"}
// @Param data body request.RegisterReq{} true "添加请求参数"
// @Router  /user/v1/register  [post]
// @version 1.0
func Register(c *gin.Context) {
	req := request.RegisterReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		user.RegisterAPI.Register(c, req)
	}
}
