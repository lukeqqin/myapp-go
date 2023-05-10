package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"Code"`
	Msg  string      `json:"Msg"`
	Data interface{} `json:"Data"`
}

// Success 封装成功返回的对象
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}

// Fail 封装失败返回的对象
func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Result{
		Code: http.StatusInternalServerError,
		Msg:  msg,
		Data: nil,
	})
}
