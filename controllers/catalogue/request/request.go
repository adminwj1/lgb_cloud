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
type CatalogueCreateRes struct {
	CatlogueInfo CatalogueInfo
}

type DelCatalogues struct {
	BUcketName    string `form:"bucket_name" binding:"required"`
	BucketID      int64  `form:"bucket_id" binding:"required"`
	CatalogueName string `form:"catalogue_name" binding:"required"`
	CatalogueId   int64  `form:"catalogue_id" binding:"required"`
}

type CatalogueDetailsReq struct {
	BUcketName string `form:"bucket_name" binding:"required"`
	BucketID   int64  `form:"bucket_id" binding:"required"`

	CatalogueName string `form:"catalogue_name" binding:"required"`
	CatalogueId   int64  `form:"catalogue_id" binding:"required"`
}

type CatalogueDetailsResp struct {
	List CatalogueInfo `json:"list"`
}

type CatalogueListReq struct {
	Limit int64 `form:"limit" binding:"required"`
	Page  int64 `form:"page" binding:"required"`
}

type CatalogueListRes struct {
	Count int64         `json:"count"`
	List  CatalogueInfo `json:"list"`
}

type SearchReq struct {
	CatalogueName string `json:"catalogue_name"`
	Limit         int64  `form:"limit" binding:"required"`
	Page          int64  `form:"page" binding:"required"`
}
