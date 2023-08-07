package bucket

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"clouds.lgb24kcs.cn/utils"
	"clouds.lgb24kcs.cn/utils/oss"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DelBucket struct {
}

var DelAPI DelBucket

func (d DelBucket) DelBucket(c *gin.Context, req request.DeleteBucketReq, userId int64) {
	bucket := models.Storage{}
	tx := global.APP.DB.Where("id=? and bucketname=? and userid=?", req.Id, req.BucketName, userId).First(&bucket)
	if tx.Error != nil || tx.RowsAffected == 0 {
		global.APP.Log.Error(tx.Error.Error())
		utils.Fail(c, errorx.BucketDel, "删除失败")
	} else {

		db := global.APP.DB.Delete(&bucket)
		if db.Error != nil {
			global.APP.Log.Error(tx.Error.Error())
			utils.Fail(c, errorx.BucketDel, "删除失败")
		} else {
			fmt.Println(bucket)
			err := oss.DeleteBucket(bucket.Accesskey, bucket.Secretkey, bucket.Zone, bucket.Bucketname)
			if err != nil {
				fmt.Println(err)
				utils.Fail(c, errorx.BucketDel, "删除失败")
			} else {
				utils.Success(c, "删除成功")
			}
		}

	}
}
