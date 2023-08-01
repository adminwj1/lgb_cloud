package user

import (
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"clouds.lgb24kcs.cn/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

type UserInfoAPI struct {
}

var UserInfo UserInfoAPI

func (u *UserInfoAPI) UserInfo(c *gin.Context, userId int64) {
	//userInfo := models.User{}

	var userInfo models.User
	tx := global.APP.DB.Where("id=?", userId).First(&userInfo)
	if tx.Error == nil {
		utils.Success(c, map[string]interface{}{
			"Username":  userInfo.Username,
			"Mobile":    userInfo.Mobile,
			"Create_At": userInfo.CreatedAt.Format(time.DateTime),
		})
	} else if tx.RowsAffected != 0 {
		utils.Fail(c, errorx.UserInfo, errors.New("用户数据获取失败").Error())
	} else {
		utils.Fail(c, errorx.UserInfo, tx.Error.Error())

	}
}
