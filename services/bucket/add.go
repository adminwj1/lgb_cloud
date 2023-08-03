package bucket

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"clouds.lgb24kcs.cn/utils"
	"clouds.lgb24kcs.cn/utils/oss"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type BucketAdd struct {
}

var BucketAddApi BucketAdd

func (a *BucketAdd) Add(c *gin.Context, req request.BucketAddReq, userId int64) {
	bucket := models.Storage{}
	// 判断同一个accesskey下是否已经存在同名的bucket
	tx := global.APP.DB.Where("bucketname=? AND accesskey =?", req.BucketName, req.AccessKey).First(&bucket)
	if bucket.Accesskey == req.AccessKey {
		global.APP.Log.Error(errors.New("AccessKey已被其他用户绑定").Error())
		utils.Fail(c, errorx.BucketAdd, "AccessKey已被其他用户绑定")

	} else if tx.RowsAffected != 0 {
		global.APP.Log.Error(errors.New("Bucket创建失败，改Bucket已存在").Error())
		utils.Fail(c, errorx.BucketAdd, "Bucket创建失败，改Bucket已存在")

	} else if tx.RowsAffected == 0 {
		err := oss.CreateBucket(req.AccessKey, req.SecretKey, req.Zone, req.BucketName)
		if err != nil {
			global.APP.Log.Error(err.Error())
			fmt.Println(err)
			utils.Fail(c, errorx.BucketAdd, err.Error())
		} else {
			bucket := models.Storage{
				Accesskey:  req.AccessKey,
				Alias:      req.Alias,
				Secretkey:  req.SecretKey,
				Bucketname: req.BucketName,
				Zone:       req.Zone,
				Userid:     userId,
			}
			if err := global.APP.DB.Create(&bucket).Error; err != nil {
				global.APP.Log.Error(err.Error())
				utils.Fail(c, errorx.BucketAdd, err.Error())
			} else {
				bucketInfo := request.BucketAddRes{
					BucketID:   int64(bucket.ID),
					Alias:      bucket.Alias,
					BucketName: bucket.Bucketname,
					CreateAt:   bucket.CreatedAt.Format(time.DateTime),
				}
				utils.Success(c, bucketInfo)
			}
		}
	} else {
		global.APP.Log.Error(tx.Error.Error())
		utils.Fail(c, errorx.BucketAdd, tx.Error.Error())

	}

}
