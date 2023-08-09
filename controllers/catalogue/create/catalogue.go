package create

import (
	"clouds.lgb24kcs.cn/controllers/catalogue/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/catalogue"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

// @Summary  目录对象创建
// @Success 200 {string} string json{"code","message"}
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data body request.CatalogueCreateReq{} true "添加请求参数"
// @Router  /catalogue/v1/create  [post]
// @version 1.0
func Create(c *gin.Context) {
	req := request.CatalogueCreateReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		userId := utils.CtxUserId(c)
		catalogue.Catalogue.Create(c, req, userId)
	}
}
