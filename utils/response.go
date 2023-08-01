package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 0
	FAIL    = 1
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    SUCCESS,
		"data":    data,
		"message": "OK",
	})
}

func Fail(ctx *gin.Context, Code int64, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    Code,
		"message": msg,
	})
}

