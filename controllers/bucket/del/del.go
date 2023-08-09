package del

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/bucket"
	"clouds.lgb24kcs.cn/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Summary  存储空间删除
// @Success 200 {string} string json{"code","message"}
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data body request.DeleteBucketReq{} true "删除请求参数"
// @Router /bucket/v1/del [delete]
// @version 1.0
func DELBucket(c *gin.Context) {
	req := request.DeleteBucketReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Fail(c, errorx.BucketDel, err.Error())
	} else {
		fmt.Println(req)
		userId := utils.CtxUserId(c)
		bucket.DelAPI.DelBucket(c, req, userId)
	}
}
