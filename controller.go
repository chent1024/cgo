package cgo

import "github.com/gin-gonic/gin"

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
