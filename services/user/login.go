package user

import (
	"errors"
	"time"

	"clouds.lgb24kcs.cn/controllers/user/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

type UserFun struct {
}

var UserApi UserFun

func (u *UserFun) Login(c *gin.Context, req request.LoginReq) {
	user := models.User{}
	err := global.APP.DB.Where("mobile=?", req.Mobile).First(&user).Error
	if err != nil {
		global.APP.Log.Error(err.Error())
	} else {
		if ok := utils.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); !ok {
			utils.Fail(c, errorx.UserLogin, errors.New("账号或密码错误").Error())
		} else {
			token, err := utils.CreateToken(int64(user.ID))
			now := time.Now().Unix()
			if err != nil {
				global.APP.Log.Error(err.Error())

				utils.Fail(c, errorx.UserLogin, err.Error())
			} else {
				UserInfo := request.LoginRes{
					Id:           int64(user.ID),
					Username:     user.Username,
					Token:        token,
					ExpireAt:     now + global.APP.Configuration.Server.AccessExpire,
					RefreshAfter: now + global.APP.Configuration.Server.AccessExpire/2,
				}
				utils.Success(c, UserInfo)
			}
		}

	}
}
