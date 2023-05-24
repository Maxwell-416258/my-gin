package service

import (
	"github.com/gin-gonic/gin"
	"mygin/errno"
	"mygin/model"
	"strconv"
)

func Delete(c *gin.Context) {
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendResponse(c, errno.ErrArgsNotComplete, "请给一个正确的用户id")
		return
	}
	if err := model.Delete(uint64(Id)); err != nil {
		SendResponse(c, errno.ErrDatabase, err.Error())
		return
	}
	SendResponse(c, nil, nil)
}
