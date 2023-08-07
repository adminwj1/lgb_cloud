package del

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/bucket"
	"clouds.lgb24kcs.cn/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

//type DelBucket struct {
//}
//
//var DelAPI DelBucket

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
