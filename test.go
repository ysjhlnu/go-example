package main

import (
	"go-gin-example/pkg/setting"
	"fmt"
)

//func main(){
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context){
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//	r.Run()
//}

func main() {
	fmt.Println(setting.DatabaseSetting.Name, setting.DatabaseSetting.Host, setting.DatabaseSetting.Password)
}

/*
db, err = gorm.Open("mysql",
fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//setting.DatabaseSetting.User,
"root",
"root",
//setting.DatabaseSetting.Password,
setting.DatabaseSetting.Host,
setting.DatabaseSetting.Name))

 */