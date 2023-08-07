package request

type ListBuckets struct {
	Id         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	Alias      string `json:"alias"`
	BucketName string `json:"bucketname"`
	CreateAt   string `json:"create_at"`
}
type BucketAddReq struct {
	AccessKey  string `json:"access_key" binding:"required"`
	SecretKey  string `json:"secret_key" binding:"required"`
	BucketName string `json:"bucket_name" binding:"required"`
	Zone       string `json:"zone" binding:"required"`
	Alias      string `json:"alias" binding:"required"`
}

type BucketAddRes struct {
	BucketID   int64  `json:"bucket_id"`
	Alias      string `json:"alias"`
	BucketName string `json:"bucket_name"`
	CreateAt   string `json:"create_at"`
}

// 加载当前登录用户的所有的bucket
type ListBucketReq struct {
	Limit int64 `form:"limit" binding:"required"`
	Page  int64 `form:"page" binding:"required"`
}

type ListBucketsResp struct {
	Count int64       `json:"count"`
	List  ListBuckets `json:"list"`
}

// 删除
type DeleteBucketReq struct {
	Id         int64  `form:"id" binding:"required"` // bucketId
	BucketName string `form:"bucketname" binding:"required"`
}

// bucket详细信息
type DetailBucketReq struct {
	BucketId int64 `form:"bucketId" binding:"required"`
}

type DetailBucketResp struct {
	BucketID   int64  `json:"bucket-id"`
	UserId     int64  `json:"user_id"`
	Alias      string `json:"alias"`
	AccessKey  string `json:"access_key"`
	SecretKey  string `json:"secret_key"`
	BucketName string `json:"bucket_name"`
	Zone       string `json:"zone"`
	CreateAt   string `json:"create_at"`
}
