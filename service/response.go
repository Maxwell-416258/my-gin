package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zxmrlc/log"
	"mygin/errno"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 定义返回函数
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	//定义返回的数据
	rsp := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, rsp)
	rsp_byte, _ := json.Marshal(rsp)
	log.Debugf("[response data]:%s", string(rsp_byte))
}
