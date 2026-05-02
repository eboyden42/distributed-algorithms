package minheap

import (
	"container/heap"
)

type Evaluator interface {
	Evaluate() float64
}

type FloatHeap []Evaluator

func New() *FloatHeap {
	h := &FloatHeap{}
	heap.Init(h)
	return h
}

func (f FloatHeap) Len() int {
	return len(f)
}

func (f FloatHeap) Less(i, j int) bool {
	return f[i].Evaluate() <= f[j].Evaluate()
}

func (f *FloatHeap) Swap(i, j int) {
	temp := (*f)[i]
	(*f)[i] = (*f)[j]
	(*f)[j] = temp
}

func (f *FloatHeap) Push(x any) {
	*f = append(*f, x.(Evaluator))
}

func (f *FloatHeap) Pop() any {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[0 : n-1]
	return x
}
