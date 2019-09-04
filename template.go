package cgo

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func NewTemplate(engine *gin.Engine, c *CgoConfig) {
	cfg := Config.Tpl
	if cfg.Enable && cfg.Path != "" {
		funcs := template.FuncMap{
			"Unescaped": Unescaped,
		}

		for k, v := range c.TplFuncMap {
			funcs[k] = v
		}
		engine.SetFuncMap(funcs)
		engine.LoadHTMLGlob(Config.Tpl.Path)
	}
}

// unescape template html content
func Unescaped(x string) interface{} {
	return template.HTML(x)
}
