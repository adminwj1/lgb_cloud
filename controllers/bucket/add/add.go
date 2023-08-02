package add

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/bucket"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

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
