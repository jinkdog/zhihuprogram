package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct { //返回信息的结构体
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Ok   bool   `json:"ok"`
}

type WithData struct { //返回信息和数据的结构体
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
}

func Ok(c *gin.Context, msg string) { //自定义写入成功的信息//和上面Response结构体对应的json标签
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"ok":   true,
	})
}

func OkWithData(c *gin.Context, msg string, data interface{}) { //自定义返回成功数据的信息//和上面的WithData定义
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"ok":   true,
		"data": data,
	})
}

func Fail(c *gin.Context, code, errCode int, msg string) { //自定义返回失败的信息
	c.JSON(code, gin.H{
		"code": errCode,
		"msg":  msg,
		"ok":   false,
	})
}

func InternalErr(c *gin.Context) { //返回内部逻辑错误
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"msg":  "internal err",
		"ok":   false,
	})
}

//response用于简化提示输出反馈的内容
