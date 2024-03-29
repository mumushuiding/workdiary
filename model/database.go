package model

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/mumushuiding/workdiary/config"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

// 配置
var conf = *config.Config

// Setup 初始化一个db连接
func Setup() {
	var err error
	log.Println("数据库初始化！！")
	db, err = gorm.Open(conf.DbType, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName))
	if err != nil {
		log.Fatalf("数据库连接失败 err: %v", err)
	}
	// 启用Logger，显示详细日志
	mode, _ := strconv.ParseBool(conf.DbLogMode)

	db.LogMode(mode)
	db.SingularTable(true) //全局设置表名不可以为复数形式
	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	idle, err := strconv.Atoi(conf.DbMaxIdleConns)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(idle)
	open, err := strconv.Atoi(conf.DbMaxOpenConns)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(open)

	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").
		AutoMigrate(&ResProject{})
	db.Model(&ResProject{}).ModifyColumn("startDate", "date").ModifyColumn("endDate", "date")
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

// GetDB getdb
func GetDB() *gorm.DB {
	return db
}

// GetTx GetTx
func GetTx() *gorm.DB {
	return db.Begin()
}

// func updateTimeStampForCreateCallback(scope *gorm.Scope) {
// 	if !scope.HasError() {
// 		nowTime := time.Now()
// 		if createTimeField, ok := scope.FieldByName("Createtime"); ok {
// 			if createTimeField.IsBlank {
// 				createTimeField.Set(nowTime)
// 			}
// 		}

// 		// if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
// 		// 	if modifyTimeField.IsBlank {
// 		// 		modifyTimeField.Set(nowTime)
// 		// 	}
// 		// }
// 	}
// }
