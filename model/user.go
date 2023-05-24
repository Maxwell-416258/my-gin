package model

import (
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
	"mygin/auth"
	. "mygin/dbConnect"
)

type UserModel struct {
	Id       uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	Username string `gorm:"column:username;not null" json:"username" binding:"required"`
	Password string `gorm:"column:password;not null" json:"password" binding:"required"`
	Role     uint8  `gorm:"column:role;not null;default:0" json:"role"`
}

// 密码加密
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// 使用gin框架的增删改查
// post
func (u *UserModel) Create() error {
	return MyDB.Self.Create(&u).Error
}

// 修改密码请求参数
type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required" validate:"min=5,max=128"`
}

// put
func UpdatePassword(id uint64, password string) error {
	user := UserModel{}
	result := MyDB.Self.Model(&user).Where("id=?", id).Update("password", password)
	return result.Error
}

// delete
func Delete(id uint64) error {
	user := UserModel{}
	user.Id = id
	return MyDB.Self.Delete(&user).Error
}

// get
type QueryUer struct {
	Offset uint16 `json:"offset"`
	Limit  uint16 `json:"limit"`
}

func ListUser(offset, limit uint16) ([]*UserModel, error) {
	log.Debugf("offset is: %d,limit is: %d", offset, limit)
	users := make([]*UserModel, 0)
	if offset == 0 {
		if err := MyDB.Self.Find(&users).Error; err != nil {
			return users, err
		}
	} else {
		if limit == 0 {
			limit = viper.GetUint16("default_limit")
		}
		if err := MyDB.Self.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
			return users, err
		}
	}
	return users, nil
}
