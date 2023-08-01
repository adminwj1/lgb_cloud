package user

import (
	"clouds.lgb24kcs.cn/controllers/user/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type UserUpdataAPI struct {
}

var Updata UserUpdataAPI

func (u *UserUpdataAPI) UserUpdata(c *gin.Context, req request.UpdataReq, id int64) {
	userinfo := models.User{}
	tx := global.APP.DB.Model(&userinfo).Where("id=?", id).Update("username", req.Username)
	if tx.Error != nil {
		utils.Fail(c, errorx.UserUpload, "用户更新失败")
	} else if tx.RowsAffected == 0 {
		utils.Fail(c, errorx.UserUpload, "用户信息不存在")

	} else {
		utils.Success(c, request.UpdataResp{
			ID:       int64(userinfo.ID),
			Username: userinfo.Username,
			Mobile:   userinfo.Mobile,
			CreateAt: userinfo.CreatedAt.Format(time.DateTime),
			UpdateAt: userinfo.UpdatedAt.Format(time.DateTime),
		})
	}
}
