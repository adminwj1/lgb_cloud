package catalogue

import (
	"clouds.lgb24kcs.cn/controllers/catalogue/request"
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

type DetailObjectAPI struct {
}

var DetailAPI DetailObjectAPI

func (a *DetailObjectAPI) DetailObject(c *gin.Context, req request.CatalogueDetailsReq, UserId int64) {
	catalogueInfo := models.Catalogue{}
	tx := global.APP.DB.Where("id=? and bucketid=? and userid=?", req.CatalogueId, req.BucketID, UserId).First(&catalogueInfo)
	if tx.Error != nil || tx.RowsAffected == 0 {
		fmt.Println(tx.Error)
		global.APP.Log.Error(tx.Error.Error())
		utils.Fail(c, errorx.ObjectDetailsl, errors.New("目录详情信息获取失败").Error())
	} else {
		//查询Bucket信息，主要获取oss的秘钥信息
		BucketInfo := models.Storage{}
		first := global.APP.DB.Where("id=?", catalogueInfo.Bucketid).First(&BucketInfo)
		if first.Error != nil || first.RowsAffected == 0 {
			global.APP.Log.Error(first.Error.Error())
			utils.Fail(c, errorx.ObjectDetailsl, errors.New("目录详情信息获取失败").Error())
		} else {
			// 校验目录对象是否存在
			if err := oss.GetObject(BucketInfo.Accesskey, BucketInfo.Secretkey, BucketInfo.Zone, req.BUcketName, req.CatalogueName); err != nil {
				utils.Fail(c, errorx.ObjectDetailsl, err.Error())
			} else {
				utils.Success(c, request.CatalogueDetailsResp{List: request.CatalogueInfo{
					ID:         int64(catalogueInfo.ID),
					DiskName:   catalogueInfo.Diskname,
					BucketName: catalogueInfo.Bucketname,
					BucketId:   catalogueInfo.Bucketid,
					UserId:     catalogueInfo.Userid,
					CreateAt:   catalogueInfo.CreatedAt.Format(time.DateTime),
				}})
			}
		}

	}
}
