package request

/*目录相关操作*/
type CatalogueCreateReq struct {
	DiskName   string `json:"disk_name" binding:"required"`
	BucketName string `json:"bucket_name" binding:"required"`
	BucketId   int64  `json:"bucket_id" binding:"required"`
}

type CatalogueInfo struct {
	ID         int64  `json:"id"`
	DiskName   string `json:"disk_name"`
	BucketName string `json:"bucket_name"`
	BucketId   int64  `json:"bucket_id"`
	UserId     int64  `json:"user_id"`
	CreateAt   string `json:"create_at"`
}
type CreateRes struct {
	CatlogueInfo CatalogueInfo
}

type DelCatalogue struct {
	BucketID int64 `form:"bucket_id" binding:"required"`
	UserID   int64 `form:"user_id" binding:"required"`
}
