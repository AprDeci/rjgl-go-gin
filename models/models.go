package models

import (
	"fmt"
	"log"

	"github.com/aprdec/rjgl/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var (
		err                          error
		dbName, user, password, host string
	)

	sec := setting.Cfg.Section("database")

	dbName = sec.Key("database").MustString("rjgl")
	user = sec.Key("username").MustString("root")
	password = sec.Key("password").MustString("root")
	host = sec.Key("host").MustString("localhost")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接失败:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("获取数据库连接失败:", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	log.Print("1233213")

}
