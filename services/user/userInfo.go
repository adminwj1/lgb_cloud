package user

import (
	"errors"
	"time"

	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

type UserInfoAPI struct {
}

var UserInfo UserInfoAPI

func (u *UserInfoAPI) UserInfo(c *gin.Context, userId int64) {
	var userInfo models.User
	tx := global.APP.DB.Where("id=?", userId).First(&userInfo)
	if tx.Error == nil {
		utils.Success(c, map[string]interface{}{
			"Id":        userInfo.ID,
			"Username":  userInfo.Username,
			"Mobile":    userInfo.Mobile,
			"Create_At": userInfo.CreatedAt.Format(time.DateTime),
		})
	} else if tx.RowsAffected != 0 {
		global.APP.Log.Error(errors.New("用户数据获取失败").Error())
		utils.Fail(c, errorx.UserInfo, errors.New("用户数据获取失败").Error())
	} else {
		global.APP.Log.Error(tx.Error.Error())
		utils.Fail(c, errorx.UserInfo, tx.Error.Error())
	}
}
