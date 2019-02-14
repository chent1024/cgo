package cgo

import "html/template"

// unescape template html content
func Unescaped(x string) interface{} {
	return template.HTML(x)
}
