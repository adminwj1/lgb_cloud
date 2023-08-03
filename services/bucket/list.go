package bucket

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BucketList struct {
}

var BucketListApi BucketList

func (b BucketList) List(c *gin.Context, req request.ListBucketReq, userId int64) {

	/*用当前登录用户id查询数据库，获取当前用户所有的bucket信息*/
	buckets := []models.Storage{}
	global.APP.DB.Where("userid=?", userId).Find(&buckets)
	fmt.Println(buckets)
}
