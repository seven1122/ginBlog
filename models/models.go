package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"seven1122/ginBlog/pkg/setting"
	"time"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	AddDt int `json:"add_dt"`
	UpdateDt int `json:"update_dt"`
	DeletedDt int `json:"deleted_dt"`
}

// 用于create数据时使用的钩子
func updateTimeStampForCreateCallback(scope *gorm.Scope)  {
	if ! scope.HasError(){
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("AddDt"); ok{
			if createTimeField.IsBlank{
				createTimeField.Set(nowTime)
			}
		}
		if updateTimeField, ok := scope.FieldByName("UpdateDt"); ok {
			if updateTimeField.IsBlank{
				updateTimeField.Set(nowTime)
			}
		}
	}

}

// 用于update数据时使用的钩子
func updateTimeStampForUpdateCallback(scope *gorm.Scope)  {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateDt", time.Now().Unix())
	}

}

func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
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