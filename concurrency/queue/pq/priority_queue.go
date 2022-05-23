package pq

import (
	"container/heap"
	"errors"
)

// @Author linjiabao
// @Date   2022/5/23
// @Reference https://github.com/jupp0r/go-priority-queue

/**
ques 1：泛型传递问题
ques 2: cannot use generic type without instantiation、在泛型传递过程中、缺了一环都会报错
 	 	这次是 type itemHeap []*item 没有加泛型说明
ques 3: method(归属到struc的func) cannot have type parameters 不能用泛型
ques 4: cannot use nil as T value in return statement 泛型不能直接使用nil、需要用一个变量代替
ques 5:
*/

type PriorityQueue[T comparable] struct {
	itemHeap *itemHeap[T]
	lookup   map[T]*item[T]
}

type itemHeap[T comparable] []*item[T]

type item[T comparable] struct {
	value    T
	priority float64
	index    int
}

func New[T comparable]() PriorityQueue[T] {
	return PriorityQueue[T]{
		itemHeap: &itemHeap[T]{},
		lookup:   make(map[T]*item[T]),
	}
}

func (p *PriorityQueue[T]) Len() int {
	return p.itemHeap.Len()
}

func (p *PriorityQueue[T]) Insert(t T, priority float64) {
	//如果元素之前就存在、返回
	_, ok := p.lookup[t]
	if ok {
		return
	}

	newItem := &item[T]{
		value:    t,
		priority: priority,
		index:    0,
	}

	heap.Push(p.itemHeap, newItem)
	p.lookup[t] = newItem
}

func (p *PriorityQueue[T]) Pop() (T, error) {
	if p.itemHeap.Len() == 0 {
		var zero T //泛型不能直接使用nil、需要用一个变量代替
		return zero, errors.New("empty queue")
	}

	i := heap.Pop(p.itemHeap) //获得any类型、将它转为泛型的T
	delete(p.lookup, i.(*item[T]).value)
	return i.(*item[T]).value, nil
}

func (p *PriorityQueue[T]) UpdatePriority(t T, newPriority float64) {
	item, ok := p.lookup[t]
	//如果不在队列里、返回
	if !ok {
		return
	}

	item.priority = newPriority
	heap.Fix(p.itemHeap, item.index)

}

func (ih *itemHeap[T]) Len() int {
	return len(*ih)
}

func (ih *itemHeap[T]) Less(i, j int) bool {
	return (*ih)[i].priority < (*ih)[j].priority
}

func (ih *itemHeap[T]) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
	(*ih)[i].index = i
	(*ih)[j].index = j
}

func (ih *itemHeap[T]) Push(x any) {
	it := x.(*item[T])
	it.index = len(*ih)
	*ih = append(*ih, it)
}

func (ih *itemHeap[T]) Pop() any {
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}
