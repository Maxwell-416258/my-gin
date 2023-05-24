package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mygin/auth"
	"mygin/errno"
	"mygin/model"
	"strconv"
)

func Modify(c *gin.Context) {
	validate := validator.New()
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendResponse(c, errno.ErrArgsNotComplete, "请输入正确的id")
		return
	}
	var passwordRequest model.UpdatePasswordRequest
	if err := c.BindJSON(&passwordRequest); err != nil {
		SendResponse(c, errno.Errbind, err.Error())
		return
	}
	if err := validate.Struct(passwordRequest); err != nil {
		SendResponse(c, errno.Errbind, err.Error())
	}

	hashedpassword, err := auth.Encrypt(passwordRequest.Password)
	if err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	//更新密码数据
	if err := model.UpdatePassword(uint64(Id), hashedpassword); err != nil {
		SendResponse(c, errno.ErrDatabase, err.Error())
		return
	}
	// 返回创建的用户信息
	SendResponse(c, nil, nil)
}
