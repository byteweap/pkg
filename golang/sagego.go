package golang

import (
	"runtime"

	"github.com/byteweap/pkg/logs"
)

// Recover panic异常堆栈信息
func Recover() {
	if err := recover(); err != nil {
		logs.Error().Msgf("panic recover err: %v", err)
		for i := 0; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			logs.Error().Msgf("%s: %d", file, line)
		}
	}
}

// SafeGo executes the given function in a separate goroutine, recovering from any panics.
func SafeGo(fn func()) {
	go func() {
		defer Recover()
		// do
		fn()
	}()
}
