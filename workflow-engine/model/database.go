package model

import (
	"fmt"
	"log"

	config "github.com/mumushuiding/go-workflow/workflow-config"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

// Model 其它数据结构的公共部分
type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

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
	db.LogMode(true)

	db.SingularTable(true) //全局设置表名不可以为复数形式
	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.DB().SetMaxIdleConns(conf.DbMaxIdleConns)
	db.DB().SetMaxOpenConns(conf.DbMaxOpenConns)

	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&Procdef{}).
		AutoMigrate(&Execution{}).AutoMigrate(&Task{}).
		AutoMigrate(&ProcInst{}).AutoMigrate(&Identitylink{})
	db.Model(&Procdef{}).AddIndex("idx_id", "id")
	db.Model(&ProcInst{}).AddIndex("idx_id", "id")
	db.Model(&Execution{}).AddForeignKey("proc_inst_id", "proc_inst(id)", "CASCADE", "RESTRICT").AddIndex("idx_id", "id")
	db.Model(&Identitylink{}).AddForeignKey("proc_inst_id", "proc_inst(id)", "CASCADE", "RESTRICT").AddIndex("idx_id", "id")
	db.Model(&Task{}).AddForeignKey("proc_inst_id", "proc_inst(id)", "CASCADE", "RESTRICT").AddIndex("idx_id", "id")
	// db.Model(&Comment{}).AddForeignKey("proc_inst_id", "proc_inst(id)", "CASCADE", "RESTRICT")
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
