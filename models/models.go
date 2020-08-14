package models

import (
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-gin-example/pkg/setting"
)

// https://eddycjy.com/posts/go/gin/2018-02-11-api-01/


// models的初始化使用

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}


func init(){
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil{
		log.Fatalf(2,"Fail to get section 'database':%v", err)
	}
}