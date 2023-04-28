package res

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ResponseList[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

const (
	Success = 0
	Error   = -100
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(Success, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func ResultList(list any, count int64, c *gin.Context) {
	OKWithData(ResponseList[any]{
		List:  list,
		Count: count,
	}, c)
}

func OK(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OKWithList(list any, count int64, c *gin.Context) {
	ResultList(list, count, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(Success, data, "成功！", c)
}

func OKWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}

func FailWithData(data any, c *gin.Context) {
	Result(Error, data, "成功！", c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	errMap := InitError()
	msg, ok := errMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(Error, map[string]any{}, "未知错误", c)
}
