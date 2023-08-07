package create

import (
	"clouds.lgb24kcs.cn/controllers/catalogue/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	req := request.CatalogueCreateReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
	}
}
