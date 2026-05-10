package main

import (
	"github.com/eboyden42/distributed-algorithms/cmd/internal/distancevector"
	"github.com/eboyden42/distributed-algorithms/cmd/internal/graph"
)

type Message struct {
	time float64
	data string
}

func (m Message) Evaluate() float64 {
	return m.time
}
func main() {
	adj := [][]float64{
		{0, 2, 0, 2, 0},
		{2, 0, 8, 0, 0},
		{0, 8, 0, 2, 1},
		{2, 0, 2, 0, 4},
		{0, 0, 1, 4, 0},
	}
	g, err := graph.NewFloatAdj(adj)
	if err != nil {
		panic(err)
	}
	dv := distancevector.New(g)
	dv.Run()
	dv.PrintRoutingTables()
}
