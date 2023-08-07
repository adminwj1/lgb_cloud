package bucket

import (
	"clouds.lgb24kcs.cn/controllers/bucket/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"clouds.lgb24kcs.cn/utils"
	"clouds.lgb24kcs.cn/utils/oss"
	"github.com/gin-gonic/gin"
	"time"
)

type DetailApi struct {
}

var Detail DetailApi

func (a DetailApi) DetailBucket(c *gin.Context, req request.DetailBucketReq, UserId int64) {
	bucketInfo := models.Storage{}
	tx := global.APP.DB.Where("id=? and userid=?", req.BucketId, UserId).First(&bucketInfo)
	if tx.Error != nil {
		global.APP.Log.Error(tx.Error.Error())
		utils.Fail(c, errorx.BucketDetaillogic, "没有数据")
	} else if tx.RowsAffected == 0 {
		utils.Fail(c, errorx.BucketDetaillogic, "没有数据")
	} else {
		ok := oss.GetBucket(bucketInfo.Accesskey, bucketInfo.Secretkey, bucketInfo.Zone, bucketInfo.Bucketname)
		if ok {
			utils.Success(c, request.DetailBucketResp{
				BucketID:   int64(bucketInfo.ID),
				UserId:     bucketInfo.Userid,
				Alias:      bucketInfo.Alias,
				AccessKey:  bucketInfo.Accesskey,
				SecretKey:  bucketInfo.Secretkey,
				BucketName: bucketInfo.Bucketname,
				Zone:       bucketInfo.Zone,
				CreateAt:   bucketInfo.CreatedAt.Format(time.DateTime),
			})
		} else {
			utils.Fail(c, errorx.BucketDetaillogic, "没有数据")
		}

	}
}
