package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// https://eddycjy.com/posts/go/gin/2018-02-11-api-01/


// models的初始化使用

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn int `json:"deleted_on"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	result := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			//setting.DatabaseSetting.User,
			"root",
			"root",
			//setting.DatabaseSetting.Password,
			//setting.DatabaseSetting.Host,
			"127.0.0.1",
			//setting.DatabaseSetting.Name)
			"blog")

	fmt.Println("eeeee:", result)
	db, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		//setting.DatabaseSetting.User,
		"root",
		"root",
		//setting.DatabaseSetting.Password,
		"127.0.0.1:3306",
		"blog"))

	if err != nil {
		//logging.Error(err)
		fmt.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		//return setting.DatabaseSetting.TablePrefix + defaultTableName
		return "blog_" + defaultTableName
	}

	db.SingularTable(true)
	// 注册 Callbacks
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:upate_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

/*
func init(){
	var (
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil{
		logging.Info(2,"Fail to get section 'database':%v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err  = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil{
		logging.Info(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 注册 Callbacks
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:upate_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
}

 */

// 关闭数据库
func CloseDB() {
	defer db.Close()
}

// 实现Callbacks
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}


func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// scope.Get(...) 根据入参获取设置了字面值的参数，例如本文中是 gorm:update_column ，它会去查找含这个字面值的字段属性
	if _, ok := scope.Get("gorm:update_column"); !ok{
		// scope.SetColumn(...) 假设没有指定 update_column 的字段，我们默认在更新回调设置 ModifiedOn 的值
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// scope.Get("gorm:delete_option") 检查是否手动指定了 delete_option
		if str,ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		// scope.FieldByName("DeletedOn") 获取我们约定的删除字段，若存在则 UPDATE 软删除，若不存在则 DELETE 硬删除
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(), 	// scope.QuotedTableName() 返回引用的表名，这个方法 GORM 会根据自身逻辑对表名进行一些处理
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),	 // scope.AddToVars 该方法可以添加值作为 SQL 的参数，也可用于防范 SQL 注入
				addExtraSpaceIfExist(scope.CombinedConditionSql()),	// scope.CombinedConditionSql() 返回组合好的条件 SQL，看一下方法原型很明了
				addExtraSpaceIfExist(extraOption),
				)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
				)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string{
	if str != "" {
		return " " + str
	}
	return ""
}