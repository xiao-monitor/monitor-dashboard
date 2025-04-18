package base

// BaseResponse 基础响应
//
//	@param Code int 状态码
//	@param Message string 消息
//	@param Data interface{} 数据(可选)
type BaseResponse struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"操作成功"`
	Data    interface{} `json:"data,omitempty" example:"{}"`
}
