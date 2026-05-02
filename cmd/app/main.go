package main

import (
	"container/heap"
	"fmt"

	"github.com/eboyden42/distributed-algorithms/cmd/internal/minheap"
)

type Message struct {
	time float64
	data string
}

func (m Message) Evaluate() float64 {
	return m.time
}
func main() {
	// adj := [][]float64{
	// 	{0, 1, 0, 1, 0},
	// 	{1, 0, 0, 0, 0},
	// 	{0, 0, 0, 1, 0},
	// 	{1, 0, 1, 0, 1},
	// 	{0, 0, 0, 1, 0},
	// }
	// g, err := graph.New(adj)
	// if err != nil {
	// 	panic(err)
	// }
	// s := sptree.New(g)
	// s.Run()

	mh := &minheap.FloatHeap{}
	heap.Init(mh)

	m1 := &Message{0.0, "hello"}
	m2 := &Message{0.5, "I was sending!"}
	m3 := &Message{0.25, "this is the message"}

	heap.Push(mh, m1)
	heap.Push(mh, m2)
	heap.Push(mh, m3)

	for mh.Len() > 0 {
		fmt.Println(heap.Pop(mh))
	}

}
