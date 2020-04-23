package e


var msgFlags = map[int]string{
	Success: "ok",
	Error: "fail",
	InvalidParams: "参数错误",
	ErrorNotExistComment: "留言不存在",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已经超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "认证错误",
	ErrorLogin: "用户名或密码错误",
	ErrorRegisterUser: "注册用户错误",
	ErrorRegisterUserExist: "注册用户已经存在",
}

func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok{
		return msg
	}
	return msgFlags[Error]
}
