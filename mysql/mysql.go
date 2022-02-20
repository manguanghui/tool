// Package mysql
/**
  * @Author:manguanghui
  * @Date: 2021/11/22
  * @Desc:
**/
package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

import (
	"log"
	"strings"
)

var db *gorm.DB

type DBConf struct {
	UserName  string
	Password  string
	IP        string
	Port      string
	DbName    string
	OpenCount int
	IdleCount int
}

// LoadMysqlDb 加载数据库驱动
func LoadMysqlDb(conf *DBConf) {
	path := strings.Join([]string{conf.UserName, ":", conf.Password, "@tcp(", conf.IP, ":", conf.Port, ")/", conf.DbName, "?charset=utf8"}, "")
	dB, err := gorm.Open("mysql", path)
	if err != nil {
		log.Fatal("出现错误：", err)
	}
	db = dB
	// 将数据库链接存入连接池
	db.DB().SetMaxOpenConns(conf.OpenCount)
	db.DB().SetMaxIdleConns(conf.IdleCount)
	err = db.DB().Ping()
	if err != nil {
		log.Fatal("出现错误：", err)
	}
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return db
}
