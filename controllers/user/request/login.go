package request

type LoginReq struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	Token        string `json:"token"`
	ExpireAt     int64  `json:"expireAt"`
	RefreshAfter int64  `json:"refreshAfter"`
}
