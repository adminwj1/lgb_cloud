package bucket

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DetailApi struct {
}

var Detail DetailApi

func (a DetailApi) DetailBucket(c *gin.Context, req request.DetailBucketReq, UserId int64) {
	bucketInfo := models.Storage{}
	fmt.Println(UserId, req.BucketId)
	tx := global.APP.DB.Where("id=?", req.BucketId).First(&bucketInfo)
	if tx != nil {
		fmt.Println(tx.Error)
	} else if tx.RowsAffected == 0 {

	} else {
		fmt.Println(bucketInfo)

	}
	//oss.GetBucket(bucketInfo.Accesskey, bucketInfo.Secretkey, bucketInfo.Zone, bucketInfo.Bucketname)
}
