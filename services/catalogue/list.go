package catalogue

import (
	"clouds.lgb24kcs.cn/controllers/catalogue/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

func CatalogueList(c *gin.Context) {
	req := request.CatalogueListReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Fail(c, errorx.ObjectList, err.Error())
	} else {
	}
}
