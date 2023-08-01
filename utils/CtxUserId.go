package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// (int64, error)
func CtxUserId(ctx *gin.Context) int64 {
	value := ctx.Value("userId")
	uid, _ := strconv.Atoi(value.(string))
	return int64(uid)

}
