package bucket

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"clouds.lgb24kcs.cn/utils"
	"clouds.lgb24kcs.cn/utils/oss"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type BucketAdd struct {
}

var BucketAddApi BucketAdd

func (a *BucketAdd) Add(c *gin.Context, req request.BucketAddReq, userId int64) {
	bucket := models.Storage{}

	if err := oss.CreateBucket(req.AccessKey, req.SecretKey, req.Zone, req.BucketName); err != nil {
		utils.Fail(c, errorx.BucketAdd, err.Error())
	} else {
		// 判断数据库是否存在同一条数据，用户id和bucketname进行判断
		tx := global.APP.DB.Where("bucketname=? and userid=?", req.BucketName, userId).First(&bucket)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			bucketInfo := models.Storage{
				Accesskey:  req.AccessKey,
				Alias:      req.Alias,
				Secretkey:  req.SecretKey,
				Bucketname: req.BucketName,
				Zone:       req.Zone,
				Userid:     userId,
			}
			// 如果数据库创建失败，需要删除bucket数据否则会导致数据不一致
			create := global.APP.DB.Create(&bucketInfo)
			if create.Error != nil {
				// 创建失败删除bucket保持数据一致性
				if err := oss.DeleteBucket(req.AccessKey, req.SecretKey, req.Zone, req.BucketName); err != nil {
					global.APP.Log.Error(err.Error())
					utils.Fail(c, errorx.BucketAdd, "Bucket创建失败")
				}
			} else {
				utils.Success(c, request.BucketAddRes{
					BucketID:   int64(bucketInfo.ID),
					Alias:      bucketInfo.Alias,
					BucketName: bucketInfo.Bucketname,
					CreateAt:   bucketInfo.CreatedAt.Format(time.DateTime),
				})
			}
		} else {
			//	数据存在
			utils.Fail(c, errorx.BucketAdd, "改Bucket已存在")
		}
	}

}
