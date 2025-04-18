package enum

import "strconv"

// ErrorCode 错误码
//
//	@param Code int 错误码
//	@param Message string 错误消息
//	@param Data interface{} 错误数据
type ErrorCode struct {
	Code    int
	Message string
	Data    interface{}
}

// 错误码映射
var errorCodeMap map[int]ErrorCode

// 初始化错误码映射
func init() {
	errorCodeMap = map[int]ErrorCode{
		// 系统错误 (500xx)
		50000: {Code: 50000, Message: "服务器内部错误", Data: nil},
		50300: {Code: 50300, Message: "服务暂时不可用", Data: nil},
		50401: {Code: 50401, Message: "数据库操作超时", Data: nil},
		50402: {Code: 50402, Message: "第三方服务调用失败", Data: nil},

		// 400 系列错误（客户端请求错误）
		40000: {Code: 40000, Message: "请求格式错误", Data: nil},
		40001: {Code: 40001, Message: "JSON解析失败", Data: nil},
		40002: {Code: 40002, Message: "XML解析失败", Data: nil},
		40010: {Code: 40010, Message: "缺少必要参数", Data: nil},
		40011: {Code: 40011, Message: "参数类型错误", Data: nil},
		40012: {Code: 40012, Message: "参数超出允许范围", Data: nil},
		40013: {Code: 40013, Message: "日期格式无效（应使用YYYY-MM-DD）", Data: nil},
		40014: {Code: 40014, Message: "时间格式无效（应使用YYYY-MM-DD HH:mm:ss）", Data: nil},
		40020: {Code: 40020, Message: "文件格式不支持", Data: nil},
		40021: {Code: 40021, Message: "文件大小超过限制", Data: nil},

		// 401 系列错误（认证错误）
		40100: {Code: 40100, Message: "未授权访问", Data: nil},
		40101: {Code: 40101, Message: "用户名或密码错误", Data: nil},
		40102: {Code: 40102, Message: "无效的认证令牌", Data: nil},
		40103: {Code: 40103, Message: "认证令牌已过期", Data: nil},
		40104: {Code: 40104, Message: "请求签名无效", Data: nil},

		// 403 系列错误（权限错误）
		40300: {Code: 40300, Message: "禁止访问", Data: nil},
		40301: {Code: 40301, Message: "权限不足", Data: nil},
		40302: {Code: 40302, Message: "IP地址被禁止访问", Data: nil},
		40303: {Code: 40303, Message: "API访问权限未开通", Data: nil},

		// 404 系列错误
		40400: {Code: 40400, Message: "资源未找到", Data: nil},
		40401: {Code: 40401, Message: "API端点不存在", Data: nil},

		// 409 系列错误（资源冲突）
		40900: {Code: 40900, Message: "资源状态冲突", Data: nil},
		40901: {Code: 40901, Message: "重复条目", Data: nil},

		// 422 系列错误（业务验证错误）
		42200: {Code: 42200, Message: "数据验证失败", Data: nil},
		42201: {Code: 42201, Message: "邮箱格式无效", Data: nil},
		42202: {Code: 42202, Message: "手机号格式无效", Data: nil},
		42203: {Code: 42203, Message: "密码必须包含大小写字母和数字（8-20位）", Data: nil},
		42204: {Code: 42204, Message: "URL格式无效", Data: nil},
		42205: {Code: 42205, Message: "IP地址格式无效", Data: nil},
		42206: {Code: 42206, Message: "值不在允许的范围内", Data: nil},

		// 429 系列错误（限流）
		42900: {Code: 42900, Message: "请求过于频繁", Data: nil},
		42901: {Code: 42901, Message: "API调用频率超限", Data: nil},
	}
}

// GetErrorCode 根据错误码获取对应的ErrorCode
//
//	@param code int 错误码
//	@return ErrorCode 错误码实例
func GetErrorCode(code int) ErrorCode {
	if errorCode, ok := errorCodeMap[code]; ok {
		return errorCode
	}
	// 未找到对应的错误码，返回默认错误
	return AddErrorCodeData(SystemInternalErrorCode, map[string]interface{}{
		"error_message": "[" + strconv.Itoa(code) + "] 错误码无法对应",
		"error_type":    "开发者开发错误",
	})
}

// AddCustomErrorCode 添加自定义错误码
//
//	@param code int 错误码
//	@param message string 错误消息
//	@param data interface{} 错误数据
//	@return ErrorCode 错误码
func AddCustomErrorCode(code int, message string, data interface{}) ErrorCode {
	errorCode := ErrorCode{Code: code, Message: message, Data: data}
	return errorCode
}

// AddErrorCodeData 添加错误码数据
//
//	@param code int 错误码
//	@param data interface{} 错误数据
//	@return ErrorCode 错误码
func AddErrorCodeData(code int, data interface{}) ErrorCode {
	if errorCode, ok := errorCodeMap[code]; ok {
		errorCode.Data = data
		return errorCode
	}
	return AddErrorCodeData(SystemInternalErrorCode, map[string]interface{}{
		"error_message": "[" + strconv.Itoa(code) + "] 错误码无法对应",
		"error_type":    "开发者开发错误",
	})
}

// 常用错误码常量，用于方便引用
const (
	SystemInternalErrorCode     = 50000
	ServiceUnavailableCode      = 50300
	DatabaseTimeoutCode         = 50401
	ThirdPartyAPIErrorCode      = 50402
	BadRequestCode              = 40000
	InvalidJSONCode             = 40001
	InvalidXMLCode              = 40002
	MissingRequiredParamCode    = 40010
	UnauthorizedCode            = 40100
	InvalidCredentialsCode      = 40101
	InvalidTokenCode            = 40102
	TokenExpiredCode            = 40103
	InvalidSignatureCode        = 40104
	ForbiddenCode               = 40300
	InsufficientPermissionsCode = 40301
	IPAddressBlockedCode        = 40302
	APIUnauthorizedCode         = 40303
	ResourceConflictCode        = 40900
	DuplicateEntryCode          = 40901
	DataValidationFailedCode    = 42200
	InvalidEmailFormatCode      = 42201
	InvalidPhoneFormatCode      = 42202
	InvalidPasswordFormatCode   = 42203
	InvalidURLFormatCode        = 42204
	InvalidIPAddressFormatCode  = 42205
	RateLimitExceededCode       = 42900
	APIThrottlingCode           = 42901
)
