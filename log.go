package cgo

import (
	"fmt"
	"os"
	"strings"
)

func Log(format string, values ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Fprintf(os.Stderr, "[Cgo-debug] "+format, values...)
}
