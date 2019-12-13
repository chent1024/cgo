package util

import (
	"fmt"

	"github.com/chent1024/cgo"
)

// 生成缓存key值
func MakeCacheKey(key string) string {
	return fmt.Sprintf("%s_%s", cgo.Config.App.Name, key)
}
