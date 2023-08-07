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

type BucketList struct {
}

var BucketListApi BucketList

func (b BucketList) List(c *gin.Context, req request.ListBucketReq, userId int64) {

	/*用当前登录用户id查询数据库，获取当前用户所有的bucket信息*/
	buckets := []models.Storage{}
	var Count int64
	tx := global.APP.DB.Where("userid=?", userId).Find(&buckets).Count(&Count)
	if tx.Error != nil {
		global.APP.Log.Error(tx.Error.Error())
		utils.Fail(c, errorx.BucketList, tx.Error.Error())
	} else if tx.RowsAffected == 0 {
		utils.Fail(c, errorx.BucketList, errors.New("没有数据").Error())
	} else {
		// 检查bucket是否存在oss存储中
		list := []request.ListBucketsResp{}
		for _, item := range buckets {
			sve := oss.NewAws(item.Accesskey, item.Secretkey, item.Zone)
			if exists := oss.BucketExists(sve, item.Bucketname); !exists {
				fmt.Println(item)

			} else {
				list = append(list, request.ListBucketsResp{
					Count: Count,
					List: request.ListBuckets{
						Id:         int64(item.ID),
						UserID:     item.Userid,
						Alias:      item.Alias,
						BucketName: item.Bucketname,
						CreateAt:   item.CreatedAt.Format(time.DateTime),
					},
				})
			}
		}
		utils.Success(c, list)
	}

}
