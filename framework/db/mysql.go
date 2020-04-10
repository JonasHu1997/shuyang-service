package db

import (
	"database/sql"
	"fmt"
	"gin-service/framework/conf"
	"log"
	"time"

	// 导入以调用它的init方法
	_ "github.com/go-sql-driver/mysql"
)

var (
	// DB 数据库实例
	DB *sql.DB
)

// InitMySQL 从配置中初始化数据库
func InitMySQL(cnf conf.IMySQL) (err error) {
	conf := conf.MySQLDefault()
	err = cnf.BindMySQL(conf)
	if err != nil {
		return
	}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=10s",
		conf.User,
		conf.Passwd,
		conf.Host,
		conf.Port,
		conf.Dbname,
		conf.Charset)
	log.Println("Mysql config:", dataSourceName)
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return
	}
	DB.SetMaxIdleConns(2)
	DB.SetMaxOpenConns(conf.MaxOpenConns)
	DB.SetConnMaxLifetime(time.Duration(60) * time.Second)
	log.Println("Mysql init Done!")
	return
}
