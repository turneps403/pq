// Package prqueue implements a priority queue with generics and ability to send comparator-func as a constructor argument.
//
// Synopsis:
//
// 	pqmin := prqueue.New(func(a, b int) bool { return a < b })
// 	pqmax := prqueue.New(func(a, b int) bool { return a > b })
//
// 	type custom struct {
// 		w int
// 	}
// 	pq := prqueue.New(func(a, b custom) bool { return a.w < b.w })
//
// 	pq.Add(el)
// 	el, err := pq.Poll()	// Retrieves and removes
// 	el, err := pq.Peek()	// Retrieves, but does not remove
// 	pq.IsEmpty()			// bool
// 	pq.Len()				// bool
package prqueue

import (
	"container/heap"
	"errors"
	"fmt"
	"sync"
)

// ErrEmptyQueue is error which indicates about empty queue
var ErrEmptyQueue = errors.New("priority queue is empty")

type PQ[T any] struct {
	mu   sync.RWMutex
	list []T
	cmp  func(a, b T) bool
}

func New[T any](cmp func(a, b T) bool, c ...int) (pq *PQ[T]) {
	pq = &PQ[T]{cmp: cmp}
	if len(c) != 0 {
		pq.list = make([]T, 0, c[0])
	} else {
		pq.list = make([]T, 0)
	}
	return
}

func (pq *PQ[T]) Add(el T) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	heap.Push(pq, el)
}

func (pq *PQ[T]) Poll() (el T, er error) {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	if len(pq.list) > 0 {
		el = heap.Pop(pq).(T)
		return
	}
	er = ErrEmptyQueue
	return
}

func (pq *PQ[T]) Peek() (el T, er error) {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	if len(pq.list) > 0 {
		el = pq.list[0]
		return
	}
	er = ErrEmptyQueue
	return
}

func (pq *PQ[T]) IsEmpty() bool {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	return len(pq.list) == 0
}

func (pq *PQ[T]) String() string {
	return fmt.Sprintf("%v", pq.list)
}

// ==========

func (pq *PQ[T]) Push(e any) {
	pq.list = append(pq.list, e.(T))
}

func (pq *PQ[T]) Len() int {
	return len(pq.list)
}

func (pq *PQ[T]) Less(i, j int) bool {
	return pq.cmp(pq.list[i], pq.list[j])
}

func (pq *PQ[T]) Swap(i, j int) {
	pq.list[i], pq.list[j] = pq.list[j], pq.list[i]
}

func (pq *PQ[T]) Pop() (e any) {
	lidx := len(pq.list) - 1
	e = pq.list[lidx]
	var tmp T
	pq.list[lidx] = tmp
	pq.list = pq.list[:lidx]
	return
}
