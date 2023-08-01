package request

type UpdataReq struct {
	Username string `json:"username" form:"username" binding:"required"`
}

type UpdataResp struct {
	ID       int64  `json:"id"`
	Username string `json:"username" `
	Mobile   string `json:"mobile"`
	CreateAt string `json:"create_at" `
	UpdateAt string `json:"update_at" `
}
