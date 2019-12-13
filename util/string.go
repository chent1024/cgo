package util

import "fmt"

// 截取字符串
func SubStr(str string, limit int, suffix string) string {
	s := []rune(str)
	var rs string
	rs = str
	if len(s) > limit {
		tmp := string(s[:limit])
		rs = fmt.Sprintf("%s%s", tmp, suffix)
	}
	return rs
}
