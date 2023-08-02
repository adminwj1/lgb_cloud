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
	// 不知道为什么oss存储对象存在相同bucketname时不会报错，这里使用查询数据库来判断bucket是否存在
	tx := global.APP.DB.Where("bucketname=? AND accesskey =?", req.BucketName, req.AccessKey).First(&bucket)
	if tx.RowsAffected != 0 {
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
