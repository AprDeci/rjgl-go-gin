package setting

import (
	"log"
	"time"

	ini "github.com/go-ini/ini"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("加载INI配置失败: %v", err)
	}
	LoadServer()

}

func LoadServer() {
	sec := Cfg.Section("app")
	RunMode = sec.Key("env").MustString("debug")
	HTTPPort = sec.Key("port").MustInt(8080)
	ReadTimeout = sec.Key("read_timeout").MustDuration(time.Second * 10)
	WriteTimeout = sec.Key("write_timeout").MustDuration(time.Second * 10)
}
