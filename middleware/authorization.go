package middleware

import (
	"regexp"
	"strings"

	"github.com/chent1024/cgo"
	"github.com/chent1024/cgo/util"

	"github.com/gin-gonic/gin"
)

// 请求头: Authorization: Bearer tokenxxx
// 解析成功后写入 context.authorization
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			authorization string
			token         []string
			path          string
			err           error
			authInfo      []byte
		)

		authorization = c.GetHeader("Authorization")
		token = strings.Split(authorization, " ")
		path = c.Request.URL.Path
		if authorization == "" || len(token) < 2 || len(token[1]) < 1 {
			if !isExceptRoute(path) {
				cgo.ResponseNoAuth(c)
				c.Abort()
				return
			}
		}

		if len(token) > 1 {
			conf := cgo.Config.Authorization
			authInfo, err = util.DesDecryptStr(token[1], conf.EncryptKey, conf.EncryptIv)
		}
		if err != nil || authInfo == nil {
			if !isExceptRoute(path) {
				cgo.ResponseNoAuth(c)
				c.Abort()
				return
			}
		}

		// 记录用户信息 方便下文使用
		c.Set("authorization", string(authInfo))

		c.Next()
	}
}

// 排除路由-不强制必须授权信息
func isExceptRoute(path string) bool {
	if len(cgo.Config.Authorization.ExpectRouter) == 0 {
		return false
	}

	for _, exp := range cgo.Config.Authorization.ExpectRouter {
		if match, _ := regexp.Match(exp, []byte(path)); match {
			return true
		}
	}

	return false
}
