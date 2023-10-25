package limit

import (
	"fmt"
	"time"

	"github.com/byteweap/pkg/golang"
	"github.com/byteweap/pkg/logs"
)

// ApiLimiter Api接口调用限速器
// 合并多个api请求为批量处理请求,需要接口方支持批量处理
// 例: 300毫秒内
type ApiLimiter struct {
	data     chan LimiterOption // 数据通道
	close    chan struct{}      // 关闭通道
	interval int64              // 时间间隔(单位:毫秒) - 严格控制此数值的修改,值越大延迟越高
}

type LimiterOption struct {
	Data     any      // Api请求数据
	RespChan chan any // 响应通道
}

func New(interval int64) *ApiLimiter {
	return &ApiLimiter{
		data:     make(chan LimiterOption),
		close:    make(chan struct{}),
		interval: interval,
	}
}

// 主程序运行
func (l *ApiLimiter) Run(fn func(datas []any) []any) {

	golang.SafeGo(func() {

		ds := make([]any, 0)         // 数据
		resps := make([]chan any, 0) // 响应信道
		lastTime := time.Now().UnixMilli()
		for {
			select {
			case data := <-l.data:
				ds = append(ds, data.Data)
				resps = append(resps, data.RespChan)

				curTime := time.Now().UnixMilli() // 毫秒
				if curTime-lastTime > l.interval {
					do(ds, resps, fn)
					ds, resps = nil, nil
					lastTime = curTime
				}
			case <-time.After(time.Millisecond * time.Duration(l.interval)):
				logs.Info(" -------------- do one !!")
				if len(ds) > 0 {
					do(ds, resps, fn)
					ds, resps = nil, nil
					lastTime = time.Now().UnixMilli() // 毫秒
				}
			case <-l.close:
				fmt.Println("close run")
				close(l.close)
				return
			}
		}
	})

}

// DO
func do(datas []any, resps []chan any, fn func([]any) []any) {
	results := fn(datas)

	if len(results) == len(resps) {
		for i := 0; i < len(results); i++ {
			resps[i] <- results[i]
		}
	} else {

	}

}

func (l *ApiLimiter) Close() {
	l.close <- struct{}{}
}

// Action
// desc: 添加要处理的数据
func (l *ApiLimiter) Action(data interface{}) chan any {

	respChan := make(chan any)
	l.data <- LimiterOption{
		Data:     data,
		RespChan: respChan,
	}
	return respChan
}
