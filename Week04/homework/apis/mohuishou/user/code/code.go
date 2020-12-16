package code

import "github.com/mohuishou/go-training/Week04/homework/errors"

// TODO: 自动生成
var (
	// UserNotFound 用户不存在
	UserNotFound = errors.NotFound("mohuishou.user.UserNotFound", "用户不存在")

	Unknown = errors.Internal("mohuishou.user.Unknown", "未知错误")
)
