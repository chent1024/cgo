package cgo

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type JsonData struct {
	Code int         `json:"code""`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 统一json输出格式
func JsonMsg(code int, obj interface{}, msg string) *JsonData {
	if msg == "" {
		msg = http.StatusText(code)
	}

	return &JsonData{
		Code: code,
		Msg:  msg,
		Data: obj,
	}
}

// Load template
func View(ctx *gin.Context, code int, tpl string, obj interface{}) {
	if code == http.StatusNotFound && tpl == "" {
		ctx.String(404, "Page Not Found")
		return
	}

	if strings.HasSuffix(tpl, "html") == false {
		tpl += ".html"
	}

	if code == 0 {
		code = http.StatusOK
	}
	ctx.HTML(code, tpl, obj)
}

func View404(ctx *gin.Context, tpl string) {
	View(ctx, http.StatusNotFound, tpl, nil)
}
