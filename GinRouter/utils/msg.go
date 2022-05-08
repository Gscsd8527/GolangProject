package utils


//不同响应码对应不同的信息
var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	ERROR_GET_USER_FAIL : "获取对象失败",
	ERROR_CREATE_USER_FAIL : "创建用户失败",
	ERROR_USER_EXIST_FAIL : "用户已经存在",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}


//获取参数
func GetMsg(code int) string{
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}