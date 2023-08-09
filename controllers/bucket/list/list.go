package list

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/bucket"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 存储空间列表
// @Success 200 {string} string json{"code","message"}
// @Param limit   query int  true  "当前页数"
// @Param page   query int  true  "查询总页数"
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Router  /bucket/v1/list  [get]
// @version 1.0
func List(c *gin.Context) {
	req := request.ListBucketReq{}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		userId := utils.CtxUserId(c)
		bucket.BucketListApi.List(c, req, userId)
	}
}
