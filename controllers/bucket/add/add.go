package add

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/bucket"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

// @Summary  存储空间创建
// @Success 200 {string} string json{"code","message"}
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data body request.BucketAddReq{} true "添加请求参数"
// @Router  /bucket/v1/createbucket  [post]
// @version 1.0
func Add(c *gin.Context) {
	req := request.BucketAddReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		userId := utils.CtxUserId(c)
		bucket.BucketAddApi.Add(c, req, userId)
	}
}
