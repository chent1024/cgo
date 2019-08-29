package cgo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var Db *gorm.DB

// init mysql by gorm
func NewMysql() {
	c := Config.Db
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s",
		c.Username,
		c.Password,
		c.Host,
		c.Database,
		c.Charset,
		Config.App.Timezone)

	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		Debug("db connect fail, %v", err)
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

// 查询分页数据
func Pagination(page int, pageSize int, out interface{}, orderBy string, where ...interface{}) error {
	if pageSize < 0 || page < 0 {
		return nil
	}

	var db *gorm.DB
	db = Db.Offset((page - 1) * pageSize).Limit(pageSize)
	if len(where) > 0 {
		db = db.Where(where)
	}

	if orderBy != "" {
		db = db.Order(orderBy)
	} else {
		db = db.Order("id desc")
	}

	result := db.Find(out)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}
