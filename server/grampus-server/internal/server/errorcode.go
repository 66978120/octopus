package server

type ErrorCode int

const (
	ErrorCodeInternalError           ErrorCode = 1001 // 内部错误
	ErrorCodeInvalidRequestParameter ErrorCode = 1002 // 无效参数
	ErrorAuthenticationFailed        ErrorCode = 1003 // 鉴权失败
	ErrorAPINotSupport               ErrorCode = 1004 // API不支持
)

var DefaultErrorMsg = map[ErrorCode]string{
	ErrorCodeInternalError:           "internal error",
	ErrorCodeInvalidRequestParameter: "invalid request parameter",
	ErrorAuthenticationFailed:        "authentication failed",
	ErrorAPINotSupport:               "api not support",
}
