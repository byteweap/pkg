package qo

import (
	"container/list"
	"github.com/violin8/pkg/golang"
	"sync"
)

// Qo 线性 gorouting 队列
type Qo struct {
	linearQo       *list.List
	mutexLinearQo  sync.Mutex
	mutexExecution sync.Mutex
}

// LinearQo 线性
type LinearQo struct {
	fn func()
}

// New 创建 Qo
func New() *Qo {
	q := &Qo{
		linearQo: list.New(),
	}
	return q
}

// Go 调用
func (q *Qo) Go(fn func()) {
	q.mutexLinearQo.Lock()
	q.linearQo.PushBack(&LinearQo{fn: fn})
	q.mutexLinearQo.Unlock()

	go func() {
		q.mutexExecution.Lock()
		defer q.mutexExecution.Unlock()

		q.mutexLinearQo.Lock()
		lq := q.linearQo.Remove(q.linearQo.Front()).(*LinearQo)
		q.mutexLinearQo.Unlock()

		defer golang.Recover()

		lq.fn()
	}()
}
