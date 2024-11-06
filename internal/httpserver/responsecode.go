package httpserver

import "net/http"

// UnknownError code
const (
	UnknownError = -1
	CustomError  = -2
	OK           = 0
)

// general error code
const (
	StatusParamsErr = iota + 40001
	StatusPermissionErr
	StatusSQLFilterFieldErr
)

// responseMsg map status code to msg
var responseMsg = map[int]string{
	http.StatusOK:                  "成功",
	http.StatusInternalServerError: "服务器无法响应",
	http.StatusBadRequest:          "请求参数错误",
	http.StatusUnauthorized:        "认证失败",
	http.StatusForbidden:           "禁止访问/修改",
	http.StatusNotFound:            "资源不存在",
	UnknownError:                   "未知错误",
	CustomError:                    "自定义错误",

	// toutiao error code
	OK:                      "OK",
	StatusParamsErr:         "参数错误",
	StatusPermissionErr:     "没有权限进行相关操作",
	StatusSQLFilterFieldErr: "过滤条件的field字段错误",
}

// ResponseMsg returns corresponding msg of status code
func ResponseMsg(code int) string {
	msg, ok := responseMsg[code]
	if ok {
		return msg
	}
	msg = http.StatusText(code)
	if msg != "" {
		return msg
	}

	return responseMsg[UnknownError]
}
