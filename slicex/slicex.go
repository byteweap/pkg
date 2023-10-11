package slicex

import (
	"sync"
)

// Slicex 线程安全的切片
type Slicex[T comparable] struct {
	mux  sync.RWMutex // 互斥锁
	data []T          // 源数据
}

func New[T comparable]() *Slicex[T] {
	sx := &Slicex[T]{
		mux:  sync.RWMutex{},
		data: make([]T, 0),
	}
	return sx
}

// Pop 弹出从beginIndex开始count个元素
// desc: 即(取出beginIndex开始count个元素,同时删除取出的元素)
// IMP: 要么取出count个元素,要么取出0个元素
func (sx *Slicex[T]) Pop(beginIndex, count int) []T {
	sx.mux.Lock()
	defer sx.mux.Unlock()

	if count == 0 {
		return []T{}
	}
	if beginIndex >= len(sx.data) {
		return []T{}
	}
	endIndex := beginIndex + count
	if endIndex >= len(sx.data) {
		return []T{}
	}
	// 1.取出
	data := sx.data[beginIndex:endIndex]
	// 2.删除
	sx.data = append(sx.data[:beginIndex], sx.data[endIndex:]...)

	return data
}

func (sx *Slicex[T]) Get(beginIndex, count int) {
	sx.mux.RLock()
	defer sx.mux.RUnlock()

}

func (sx *Slicex[T]) Len() int {
	sx.mux.RLock()
	defer sx.mux.RUnlock()
	return len(sx.data)
}

func (sx *Slicex[T]) Append(val T) {
	sx.mux.Lock()
	defer sx.mux.Unlock()
	sx.data = append(sx.data, val)
}

func (sx *Slicex[T]) Remove(index int) {
	sx.mux.Lock()
	defer sx.mux.Unlock()
	sx.data = append(sx.data[:index], sx.data[index+1:]...)
}

func (sx *Slicex[T]) Clear() {
	sx.mux.Lock()
	defer sx.mux.Unlock()
	sx.data = make([]T, 0)
}

func (sx *Slicex[T]) Data() []T {
	sx.mux.RLock()
	defer sx.mux.RUnlock()
	return sx.data
}
