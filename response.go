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

// 返回 成功json
func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, JsonMsg(0, data, "success"))
}

// 返回 权限校验 失败json
func ResponseNoAuth(ctx *gin.Context) {
	ctx.JSON(200, JsonMsg(-1, nil, "权限校验失败~"))
}

// 返回 参数校验 失败json
func ResponseParamErr(ctx *gin.Context) {
	ctx.JSON(200, JsonMsg(-1, nil, "参数错误~"))
}

// 返回 404错误 json
func Response404Err(ctx *gin.Context) {
	ctx.JSON(200, JsonMsg(-1, nil, "Data Not Found"))
}

// 返回 服务错误 json
func ServerErr(ctx *gin.Context, err string) {
	ctx.JSON(500, JsonMsg(-1, nil, err))
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
