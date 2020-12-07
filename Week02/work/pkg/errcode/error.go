package errcode

import "fmt"

// ErrorCode 错误码
type ErrorCode struct {
	// Status 表示 http status or grpc code
	Status int `json:"-"`
	// Code 业务错误码
	// 规则: domain.reason，例如 user.login.fail
	Code string `json:"code"`
	// Message 业务错误原因
	Message string `json:"message"`
	// Data 附加的详细信息
	Data interface{} `json:"data"`
}

// Error for error
func (e *ErrorCode) Error() string {
	return fmt.Sprintf("error: status = %d code= %s desc = %s data = %+v", e.Status, e.Code, e.Message, e.Data)
}
