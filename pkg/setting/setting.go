package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

// https://eddycjy.com/posts/go/gin/2018-02-11-api-01/

import (
	"log"
	"time"
	"github.com/go-ini/ini"
)


var (
	cfg *ini.File
	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
)

func init(){
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil{
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").key("RUN_MODE").MustString("debug")
}


func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server':%v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.key("READ_TIMEOUT").MustInt(60)) * time.Second
	WRITE_TIMEOUT = time.Duration(sec.key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSecion("app")
	if err != nil{
		log.Fatalf("Fail to get section'app':%v", err)
	}
	JwtSecret = sec.key("JWT_SECRET").MustString("!@*#)!@U#@*!@!)")
	PageSize = sec.key("PAGE_SIZE").MustInt(10)
}