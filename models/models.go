package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"seven1122/ginBlog/pkg/logging"
	"seven1122/ginBlog/pkg/setting"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	AddDt     int `json:"add_dt"`
	UpdateDt  int `json:"update_dt"`
	DeletedDt int `json:"deleted_dt"`
}

// 用于create数据时使用的钩子
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("AddDt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if updateTimeField, ok := scope.FieldByName("UpdateDt"); ok {
			if updateTimeField.IsBlank {
				updateTimeField.Set(nowTime)
			}
		}
	}

}

// 用于update数据时使用的钩子
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateDt", time.Now().Unix())
	}

}

func Setup() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = setting.DatabaseSetting.Type
	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		logging.Error("connect database err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	//配置钩子
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
}

func CloseDB() {
	defer db.Close()
}
