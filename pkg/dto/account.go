package dto

// CreateAccountReq 创建账号请求参数
type CreateAccountReq struct {
	// 用户名
	Username string `json:"username" binding:"required,min=2,max=30"`
	// 密码
	Password string `json:"password" binding:"required,min=6"`
}

// CreateAccountResp 创建账号响应
type CreateAccountResp struct {
	// 状态码
	Code int `json:"code"`
	// 消息
	Message string `json:"message"`
}
