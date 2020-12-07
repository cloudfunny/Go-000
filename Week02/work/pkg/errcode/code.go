package errcode

var (
	UserLogin = &ErrorCode{Status: 400, Code: "user.login.fail", Message: "用户名或密码错误"}

	DBQuery = &ErrorCode{Status: 500, Code: "db.query.fail", Message: "数据库查询失败"}

	ErrParams = &ErrorCode{Status: 400, Code: "err.params", Message: "参数错误"}

	ErrUnKnown = &ErrorCode{Status: 500, Code: "err.unknown", Message: "未知错误"}
)
