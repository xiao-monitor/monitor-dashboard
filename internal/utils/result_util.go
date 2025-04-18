package utils

import (
	"xiaolfeng/monitor-dashboard/internal/model/base"
	"xiaolfeng/monitor-dashboard/internal/model/enum"

	"github.com/gin-gonic/gin"
)

// Success 成功返回
//
// 作为正确返回，可以自定义返回消息和数据
//
//	@param c *gin.Context 上下文
//	@param message string 消息
//	@param data interface{} 数据
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(200, base.BaseResponse{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

// SuccessMessage 成功返回
//
// 作为正确返回，可以自定义返回消息
//
//	@param c *gin.Context 上下文
//	@param message string 消息
func SuccessMessage(c *gin.Context, message string) {
	c.JSON(200, base.BaseResponse{
		Code:    200,
		Message: message,
		Data:    nil,
	})
}

// SuccessData 成功返回
//
// 作为正确返回，可以自定义返回数据
//
//	@param c *gin.Context 上下文
//	@param data interface{} 数据
func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(200, base.BaseResponse{
		Code:    200,
		Message: "Operation successful",
		Data:    data,
	})
}

// SuccessNormal 成功返回
//
// 作为正确返回，不自定义返回消息和数据
//
//	@param c *gin.Context 上下文
func SuccessNormal(c *gin.Context) {
	c.JSON(200, base.BaseResponse{
		Code:    200,
		Message: "Operation successful",
		Data:    nil,
	})
}

// ErrorCustom 错误返回
//
// 作为错误返回，可以自定义返回消息和数据
//
//	@param c *gin.Context 上下文
//	@param errorCode int 错误码
//	@param message string 消息
//	@param data interface{} 数据
func ErrorCustom(c *gin.Context, errorCode int, message string, data interface{}) {
	c.JSON(200, base.BaseResponse{
		Code:    errorCode,
		Message: message,
		Data:    data,
	})
}

// Error 错误返回
//
// 作为错误返回，使用预定义的错误码
//
//	@param c *gin.Context 上下文
//	@param errorCode enum.ErrorCode 错误码
func Error(c *gin.Context, errorCode enum.ErrorCode) {
	c.JSON(200, base.BaseResponse{
		Code:    errorCode.Code,
		Message: errorCode.Message,
		Data:    errorCode.Data,
	})
}
