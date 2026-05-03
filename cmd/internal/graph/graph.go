package graph

import (
	"fmt"
)

type Graph [][]float64

func New(adj [][]float64) (Graph, error) {
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
	return Graph(adj), nil
}

func (g Graph) IsConnected(a, b int) bool {
	if 0 <= a && a < len(g) && 0 <= b && b < len(g[0]) {
		return g[a][b] != 0
	}
	panic(fmt.Errorf("Index (%d, %d) out of bounds for size (%d, %d). \n", a, b, len(g), len(g[0])))
}

func (g Graph) GetConnectedNodes(a int) []int {
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

func (g Graph) GetWeight(i, j int) float64 {
	return g[i][j]
}
