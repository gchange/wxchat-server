package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response struct
type ResponseResult struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg,omitempty" `
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type PageResult struct {
	Items    interface{} `json:"items"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

// Response 返回的数据
func Response(c *gin.Context, code int, msg string, data interface{}, err interface{}) {
	c.JSON(code, ResponseResult{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Error: err,
	})
}

// ResponseOK 返回成功
func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseResult{
		Code:  http.StatusOK,
		Msg:   ResponseMsg(http.StatusOK),
		Data:  data,
		Error: nil,
	})
}

// ResponseBadRequest returns 400
func ResponseBadRequest(c *gin.Context, err interface{}) {
	c.JSON(http.StatusBadRequest, ResponseResult{
		Code:  http.StatusBadRequest,
		Msg:   ResponseMsg(http.StatusBadRequest),
		Data:  nil,
		Error: err,
	})
}

// ResponseNotFound returns 404
func ResponseNotFound(c *gin.Context, err interface{}) {
	c.JSON(http.StatusNotFound, ResponseResult{
		Code:  http.StatusNotFound,
		Msg:   ResponseMsg(http.StatusNotFound),
		Data:  nil,
		Error: err,
	})
}

// ResponseUnauthorized returns 401
func ResponseUnauthorized(c *gin.Context, err interface{}) {
	c.JSON(http.StatusUnauthorized, ResponseResult{
		Code:  http.StatusUnauthorized,
		Msg:   ResponseMsg(http.StatusUnauthorized),
		Data:  nil,
		Error: err,
	})
}

// ResponseForbidden returns 403
func ResponseForbidden(c *gin.Context, err interface{}) {
	c.JSON(http.StatusForbidden, ResponseResult{
		Code:  http.StatusForbidden,
		Msg:   ResponseMsg(http.StatusForbidden),
		Data:  nil,
		Error: err,
	})
}

// ResponseServerError returns server error
func ResponseServerError(c *gin.Context, err interface{}) {
	c.JSON(http.StatusInternalServerError, ResponseResult{
		Code:  http.StatusInternalServerError,
		Msg:   ResponseMsg(http.StatusInternalServerError),
		Data:  nil,
		Error: err,
	})
}

// ResponseCustomError returns user custom error
func ResponseCustomError(c *gin.Context, code int, err interface{}) {
	c.JSON(code, ResponseResult{
		Code:  code,
		Msg:   ResponseMsg(code),
		Data:  nil,
		Error: err,
	})
}
