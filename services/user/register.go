package user

import (
	"time"

	"clouds.lgb24kcs.cn/controllers/user/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"clouds.lgb24kcs.cn/utils"
	"github.com/gin-gonic/gin"
)

type RegisterFunc struct {
}

var RegisterAPI RegisterFunc

func (r *RegisterFunc) Register(c *gin.Context, req request.RegisterReq) {
	password, _ := utils.GenerateFromPassword([]byte(req.Password))
	pwd := password
	user := models.User{
		Username: req.Username,
		Password: pwd,
		Mobile:   req.Mobile,
	}
	err := global.APP.DB.Create(&user).Error
	if err != nil {
		global.APP.Log.Error(err.Error())

		utils.Fail(c, errorx.UserRegister, err.Error())
	} else {
		userInfo := request.RegisterRes{
			Id:       int64(user.ID),
			Username: user.Username,
			Mobile:   user.Mobile,
			Create:   user.CreatedAt.Format(time.DateTime),
		}
		utils.Success(c, userInfo)
	}

}
