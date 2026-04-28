package main

import (
	"github.com/eboyden42/distributed-algorithms/cmd/internal/graph"
	"github.com/eboyden42/distributed-algorithms/cmd/internal/sptree"
)

func main() {
	adj := [][]bool{
		{false, true, false, true, false},
		{true, false, false, false, false},
		{false, false, false, true, false},
		{true, false, true, false, true},
		{false, false, false, true, false},
	}
	g, err := graph.New(adj)
	if err != nil {
		panic(err)
	}
	s := sptree.New(g)
	s.Run()
}
