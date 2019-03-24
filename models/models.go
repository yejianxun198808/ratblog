package models

import (
	"fmt"
	"log"
	"ratblog/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key"`
	CreatedOn  int
	ModifiedOn int
}

//sql数据库连接
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec := config.Config().Servers
	dbType = sec.Type
	dbName = sec.Dbname
	user = sec.Username
	password = sec.Password
	host = sec.Host
	tablePrefix = sec.Tablename

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	//更改默认表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

//数据库关闭
func CloseDB() {
	defer db.Close()
}
