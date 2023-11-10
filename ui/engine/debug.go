package engine

import (
	"log"
	"runtime"
)

// https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
func Trace() {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	log.Println("[TRACE]", line, fn.Name())
}
