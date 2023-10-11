package golang

import (
	"github.com/violin8/pkg/logs"
	"runtime"
)

//// SafeGo go函数(可捕获异常)
//func SafeGo(fn func()) {
//	go func() {
//		defer func() {
//			if err := recover(); err != nil {
//				trace := make([]byte, 1<<16)
//				n := runtime.Stack(trace, true)
//				logs.Error("SafeGo Error: %v", fmt.Errorf("panic recover\n %v\n stack trace %d bytes\n %s",
//					err, n, trace[:n]))
//			}
//		}()
//
//		// do
//		fn()
//	}()
//}

// Recover panic异常堆栈信息
func Recover() {
	if err := recover(); err != nil {
		logs.Errorx().Msgf("panic recover err: %v", err)
		for i := 0; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			logs.Errorx().Msgf("%s: %d", file, line)
		}
	}
}

// SafeGo 可捕获异常的go函数
func SafeGo(fn func()) {
	go func() {
		defer Recover()
		// do
		fn()
	}()
}
