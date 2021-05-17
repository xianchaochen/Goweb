package errno

//错误码为 5 位数
//第一位是服务级别, 1 为系统错误, 2 为普通错误.
//第二三位是模块, 模块不是指 Go 中的模块, 而是指代某个范围, 比如数据库错误, 认证错误.
//第四五位是具体错误, 比如数据库错误中的插入错误, 找不到数据等.
var (
	OK = NewError(0, "OK")

	// 服务级错误码
	ErrServer    = NewError(10001, "服务异常，请联系管理员")
	ErrParam     = NewError(10002, "参数有误")
	ErrSignParam = NewError(10003, "签名参数有误")

	// 模块级错误码 - 用户模块01
	ErrUserPhone   = NewError(20101, "用户手机号不合法")
	ErrUserCaptcha = NewError(20102, "用户验证码有误")
	ErrUserRegisterFailed = NewError(20103, "注册失败")
	ErrUserTokenEmpty = NewError(20104, "Token为空")
	ErrUserTokenInvalid = NewError(20105, "Token不合法")
	ErrUserVoteFAILED = NewError(20106, "投票失败")




)