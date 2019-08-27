package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type JsonData struct {
	Code int         `json:"code""`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response json data
func Json(ctx *gin.Context, code int, obj interface{}, msg string) {
	if msg == "" {
		msg = http.StatusText(code)
	}

	resp := &JsonData{
		Code: code,
		Msg:  msg,
		Data: obj,
	}
	ctx.JSON(code, resp)
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
