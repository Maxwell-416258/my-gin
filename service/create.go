package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zxmrlc/log"
	"mygin/errno"
	"mygin/model"
)

type CreateUserRsp struct {
	Username string
	Role     uint8
}

// 请求相应数据
func Create(c *gin.Context) {
	var user model.UserModel
	if err := c.BindJSON(&user); err != nil {
		log.Errorf(err, "注册失败，参数绑定异常")
		SendResponse(c, errno.Errbind, nil)
		return
	}
	if err := user.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	log.Debugf("user data insert %s", user.Username)
	if err := user.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, err)
		return
	}
	rsp_data := CreateUserRsp{
		Username: user.Username,
		Role:     user.Role,
	}
	SendResponse(c, nil, rsp_data)
}
