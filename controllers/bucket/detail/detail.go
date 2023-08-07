package detail

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/bucket"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

func DetailBucket(c *gin.Context) {
	req := request.DetailBucketReq{}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		userId := utils.CtxUserId(c)
		bucket.Detail.DetailBucket(c, req, userId)
	}
}
