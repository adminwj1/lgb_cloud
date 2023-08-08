package create

import (
	"clouds.lgb24kcs.cn/controllers/catalogue/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/catalogue"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	req := request.CatalogueCreateReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		userId := utils.CtxUserId(c)
		catalogue.Catalogue.Create(c, req, userId)
	}
}
