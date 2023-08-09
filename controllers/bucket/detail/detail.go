package detail

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/bucket"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 存储空间详情
// @Success 200 {string} string json{"code","message"}
// @Param        bucket_id   query      int  true  "存储空间ID"
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Router  /bucket/v1/detail  [get]
// @version 1.0
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
