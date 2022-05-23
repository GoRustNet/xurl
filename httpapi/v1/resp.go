package v1

import (
	"net/http"

	"github.com/GoRustNet/xurl/errs"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func newResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func errorToResponse(err *errs.Error) *Response {
	return newResponse(int(err.Type), err.Message, gin.H{})
}

func dataToResponse(data interface{}) *Response {
	return newResponse(0, "OK", data)
}
func resp(r *Response, c *gin.Context) {
	c.JSON(http.StatusOK, r)
}

func failed(err *errs.Error, c *gin.Context) {
	resp(errorToResponse(err), c)
}

func ok(data interface{}, c *gin.Context) {
	resp(dataToResponse(data), c)
}
