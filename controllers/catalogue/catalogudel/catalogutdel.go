package catalogudel

import (
	"clouds.lgb24kcs.cn/controllers/catalogue/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/services/catalogue"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

// @Summary  目录对象删除
// @Success 200 {string} string json{"code","message"}
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param data body request.DelCatalogues{} true "删除请求参数"
// @Router /catalogue/v1/delete [delete]
// @version 1.0
func DeleteObject(c *gin.Context) {
	req := request.DelCatalogues{}
	if err := c.ShouldBind(&req); err != nil {
		utils.Fail(c, errorx.VerifyError, err.Error())
	} else {
		userId := utils.CtxUserId(c)
		catalogue.DeleteObject.DeleteCatalogue(c, req, userId)
	}
}
