package dbConnect

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
)

type Database struct {
	Self *gorm.DB
	// 如果想连接多个数据库如Docker，这里可以再加一个数据库
	// Docker *gorm.DB
}

var MyDB *Database

func InitMyDB() *gorm.DB {
	return connectDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}
func (db *Database) Init() {
	MyDB = &Database{
		Self: GetMyDb(),
	}
}
func connectDB(username, password, addr, name string) *gorm.DB {
	dbconfig := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username, password, addr, name, true, "Local",
	)
	log.Debugf("开始连接数据库......")
	db, err := gorm.Open("mysql", dbconfig)
	if err != nil {
		log.Errorf(err, "Database connetction filed,Database name: %s", name)
	}

	//数据库设置
	setupDB(db)
	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))

	db.DB().SetMaxIdleConns(0)
}

func GetMyDb() *gorm.DB {
	return InitMyDB()
}
func (db *Database) Close() {
	MyDB.Self.Close()
}
