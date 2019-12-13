package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/chent1024/cgo/middleware"

	"github.com/chent1024/cgo"
	"github.com/chent1024/cgo/util"
	"github.com/gin-gonic/gin"
)

type Resource struct {
	ID        int             `gorm:"primary_key" json:"id"`
	CateId    int             `json:"cate_id"`
	Images    json.RawMessage `json:"images"`
	Title     string          `json:"title"`
	PageTotal int             `json:"page_total"`
	FavNum    int             `json:"fav_num"`
	FilePath  string          `json:"file_path"`
	Free      int             `json:"free"`
	Released  int             `json:"released"`
	DeletedAt *util.LocalDate `json:"deleted_at"`
	CreatedAt util.LocalDate  `json:"created_at"`
	UpdatedAt util.LocalDate  `json:"updated_at"`
}

func main() {
	config := &cgo.CgoConfig{
		ConfigPath: "app.toml",
		TplFuncMap: nil,
	}

	serv := cgo.New(config)
	serv.Use(middleware.Authorization())
	serv.GET("/", func(context *gin.Context) {
		//cgo.Db.Find(&Resource{}, "id=1")
		cgo.NewRedis()
		_, err1 := cgo.Redis.Set("aaa", 111, time.Hour).Result()
		val, err2 := cgo.Redis.Get("aaa").Result()

		cgo.Debug("", cgo.Redis, val, err1, err2)
		fmt.Fprintf(context.Writer, "hello world")
	})
	cgo.Run(serv)
}
