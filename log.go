package cgo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func Log(format string, values ...interface{}) {
	if gin.IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}

		fmt.Fprintf(os.Stderr, "[Cgo-debug] "+format, values...)
	}
}
