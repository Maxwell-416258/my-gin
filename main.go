package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
	"mygin/dbConnect"
	"mygin/parseConfig"
	"mygin/router"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
)

func main() {
	//先解析
	pflag.Parse()
	if err := parseConfig.Init(*cfg); err != nil {
		panic(err)
	}
	log.Debugf("命令行参数解析完成......")
	gin.SetMode(viper.GetString("runmode"))
	//初始化数据库
	dbConnect.MyDB.Init()
	defer dbConnect.MyDB.Close()
	//生成gin.Engine实例
	g := gin.Default()
	router.Load(g)
	g.Run(viper.GetString("addr"))

}
