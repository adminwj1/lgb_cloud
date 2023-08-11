package catalogue

import (
	"clouds.lgb24kcs.cn/controllers/catalogue/request"
	"clouds.lgb24kcs.cn/errorx"
	"clouds.lgb24kcs.cn/global"
	"clouds.lgb24kcs.cn/models"
	"clouds.lgb24kcs.cn/utils"
	"clouds.lgb24kcs.cn/utils/oss"
	"errors"
	"github.com/gin-gonic/gin"
)

type DeleteObjectAPI struct {
}

var DeleteObject DeleteObjectAPI

func (a *DeleteObjectAPI) DeleteCatalogue(c *gin.Context, req request.DelCatalogues, userId int64) {
	//查询数据oss存储对象信息获取秘钥
	BucketInfo := models.Storage{}
	tx := global.APP.DB.Where("id=? and userid=?", req.BucketID, userId).First(&BucketInfo)
	if tx.Error != nil || tx.RowsAffected == 0 {
		global.APP.Log.Error(tx.Error.Error())
		utils.Fail(c, errorx.ObjectDel, "删除失败")
	} else {
		// 查询数据库
		catalogueInfo := models.Catalogue{}
		first := global.APP.DB.Where("id=? and bucketid=? and userid=?", req.CatalogueId, req.BucketID, userId).Delete(&catalogueInfo)
		if first.Error != nil {
			utils.Fail(c, errorx.ObjectDel, first.Error.Error())
		} else if first.RowsAffected == 0 {
			global.APP.Log.Error(errors.New("数据库没有该条数据").Error())
			utils.Fail(c, errorx.ObjectDel, first.Error.Error())

		} else {
			if object := oss.DelCatalogue(BucketInfo.Accesskey, BucketInfo.Secretkey, BucketInfo.Zone, req.BUcketName, req.CatalogueName); !object {
				utils.Fail(c, errorx.ObjectDel, "删除失败")
			} else {
				utils.Success(c, "删除成功")
			}

		}
	}

}
