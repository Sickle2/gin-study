package response

var msgMap = map[ResultCode]string{
	Success:               "ok",
	Failed:                "error",
	BadRequest:            "请求参数错误",
	Unauthorized:          "认证失败",
	Forbidden:             "权限不足",
	StatusTooManyRequests: "请求过多，请稍后重试",
	NotFound:              "路径不存在",
	InternalServerError:   "服务器内部错误",
}

func GetMessage(code ResultCode) string {
	msg, ok := msgMap[code]
	if ok {
		return msg
	}

	return msgMap[Failed]
}
