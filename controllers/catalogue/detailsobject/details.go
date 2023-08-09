package detailsobject

import (
	"clouds.lgb24kcs.cn/controllers/catalogue/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/catalogue"
	"clouds.lgb24kcs.cn/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Summary 目录对象详情
// @Success 200 {string} string json{"code","message"}
// @Param bucket_name query string true  "存储空间名称"
// @Param bucket_id query int true "存储空间ID"
// @Param catalogue_id query string true "目录对象名称"
// @Param catalogue_name query int true "目录对象ID"
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Router  /catalogue/v1/detail [get]
// @version 1.0
func DetailObject(c *gin.Context) {
	req := request.CatalogueDetailsReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		fmt.Println(req)
		UserId := utils.CtxUserId(c)
		catalogue.DetailAPI.DetailObject(c, req, UserId)
	}
}
