package graph

import (
	"golang.org/x/exp/constraints"
)

// Graph types are able to have connected nodes, weighted edges, and the ability to see how many nodes are in the graph.
type Graph[T constraints.Ordered] interface {
	IsConnected(a, b int) bool
	GetConnectedNodes(a int) []int
	GetWeight(i, j int) T
	Len() int // <-- returns the number of nodes
}
