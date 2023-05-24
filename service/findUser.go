package service

import (
	"github.com/gin-gonic/gin"
	"mygin/errno"
	"mygin/model"
)

func FindUser(c *gin.Context) {
	var qu model.QueryUer
	if err := c.BindJSON(&qu); err != nil {
		SendResponse(c, errno.Errbind, err.Error())
		return
	}
	users, err := model.ListUser(qu.Offset, qu.Limit)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, err.Error())
	}
	SendResponse(c, nil, users)
}
