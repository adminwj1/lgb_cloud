package middleware

import (
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func ChecKToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			utils.Fail(c, errorx.TokenError, errors.New("token解析失败").Error())
			c.Abort()
			return
		}
		if strings.Split(tokenString, " ")[0] != "Bearer" {
			utils.Fail(c, errorx.TokenError, errors.New("token格式错误").Error())
			c.Abort()
			return

		}
		//	校验token合法性
		token, err := utils.ParseToken(strings.Split(tokenString, " ")[1])
		if err != nil {
			utils.Fail(c, errorx.TokenError, "token解析错误")
			c.Abort()
			return
		}
		// 判断过期时间
		if time.Now().Unix() > token.ExpiresAt.Unix() {
			utils.Fail(c, errorx.TokenError, "token授权过期")
			c.Abort()
			return
		} else {
			c.Set("userId", token.ID)
			c.Next()
		}
	}
}
