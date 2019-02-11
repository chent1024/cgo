package cgo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var Db *gorm.DB

// init mysql by gorm
func InitMysql() {
	c := Config.Db
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		c.Username,
		c.Password,
		c.Host,
		c.Database,
		c.Charset)

	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		Debug("db connect fail,", err)
		return
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return c.TablePrefix + defaultTableName
	}

	Db = conn
	// 开启日志
	Db.LogMode(c.Debug)
	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(c.MaxIdleConns)
	Db.DB().SetMaxOpenConns(c.MaxOpenConns)
	Db.DB().SetConnMaxLifetime(time.Second * time.Duration(c.ConnMaxLifeTime))

	Debug("init mysql connection success")
	return
}
