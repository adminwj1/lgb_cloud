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
	"time"
)

type CatalogueCreateApi struct{}

var Catalogue CatalogueCreateApi

func (catalogue *CatalogueCreateApi) Create(c *gin.Context, req request.CatalogueCreateReq, userId int64) {

	/*不能出现重名数据*/
	BucketInfo := models.Storage{}
	tx := global.APP.DB.Where("id=? AND userid=?", req.BucketId, userId).First(&BucketInfo)
	if tx.Error != nil || tx.RowsAffected == 0 {
		global.APP.Log.Error(tx.Error.Error())
		utils.Fail(c, errorx.ObjectCreate, errors.New("创建存储目录失败").Error())
	} else {
		if ok := oss.CreateObject(BucketInfo.Accesskey, BucketInfo.Secretkey, BucketInfo.Zone, req.BucketName, req.DiskName); !ok {
			utils.Fail(c, errorx.ObjectCreate, errors.New("创建存储目录失败，目录存储").Error())
		} else {
			catalogueInfo := models.Catalogue{
				Diskname:   req.DiskName,
				Userid:     userId,
				Bucketname: req.BucketName,
				Bucketid:   req.BucketId,
			}
			create := global.APP.DB.Create(&catalogueInfo)
			// 如果数据库在创建数据时出现错误，就删除oss存储中的数据
			if create.Error != nil {
				// 删除oss存储中对应的存储目录数据
				if ok := oss.DelObject(BucketInfo.Accesskey, BucketInfo.Secretkey, BucketInfo.Zone, req.BucketName, req.DiskName); !ok {
					utils.Fail(c, errorx.ObjectCreate, errors.New("创建存储目录失败").Error())
				}
			} else {
				// 如果数据库插入数据没有出现任何错误，就将创建好的数据返回给前端
				utils.Success(c, request.CatalogueCreateRes{CatlogueInfo: request.CatalogueInfo{
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
