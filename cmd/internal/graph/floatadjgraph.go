package graph

import (
	"fmt"
)

// A FloatAdjGraph is a implementation of the graph interface that uses a 2D array of floats as an adjacency matrix.
type FloatAdjGraph [][]float64

// NewFloatAdj creates a new instance of FloatAdjGraph, given a [][]float64. Returns errors if the input matrix is not square and symetric.
func NewFloatAdj(adj [][]float64) (FloatAdjGraph, error) {
	for r, l := range adj {
		// check that the 2d slice is square
		if len(l) != len(adj) {
			return nil, fmt.Errorf("Input adjacency matrix must be square. \n")
		}
		// check that the 2d slice is symetric
		for c := range len(l) {
			if adj[r][c] != adj[c][r] {
				return nil, fmt.Errorf("Input adjacency matrix must be symetric. \n")
			}
		}
	}
	return FloatAdjGraph(adj), nil
}

// IsConnected reports wether two given nodes are connected in the graph.
func (g FloatAdjGraph) IsConnected(a, b int) bool {
	if 0 <= a && a < len(g) && 0 <= b && b < len(g[0]) {
		return g[a][b] != 0
	}
	panic(fmt.Errorf("Index (%d, %d) out of bounds for size (%d, %d). \n", a, b, len(g), len(g[0])))
}

// GetConnectedNodes returns a slice of nodes connected to an input node.
func (g FloatAdjGraph) GetConnectedNodes(a int) []int {
	if 0 <= a && a < len(g) {
		res := []int{}
		for i, b := range g[a] {
			if b != 0 {
				res = append(res, i)
			}
		}
		return res
	}
	panic(fmt.Errorf("Index %d out of bounds for length %d. \n", a, len(g)))
}

// GetWeight returns the weight of a given edge between two nodes. The weight in this float graph is simply the value in the adjacency matrix.
func (g FloatAdjGraph) GetWeight(i, j int) float64 {
	return g[i][j]
}

// Len returns the number of nodes in the graph.
func (g FloatAdjGraph) Len() int {
	return len(g)
}
