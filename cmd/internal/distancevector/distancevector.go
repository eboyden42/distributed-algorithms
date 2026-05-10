package distancevector

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/eboyden42/distributed-algorithms/cmd/internal/graph"
	"github.com/eboyden42/distributed-algorithms/cmd/internal/minheap"
)

func New(g graph.Graph[float64]) *DVAlgorithm {
	nodeInfo := []Node{}
	minHeap := minheap.New()
	n := g.Len()

	// interate through each node
	for i := range n {
		// construct the inital table, and send to neighbors
		tableEntries := make([]TableEntry, n)
		neighbors := []int{}
		for node := range n {
			weight := g.GetWeight(node, i)
			if weight == 0 {
				if node != i {
					tableEntries[node].distance = math.Inf(1)
				}
			} else {
				tableEntries[node].distance = weight
				tableEntries[node].nextHop = node
				neighbors = append(neighbors, node)
			}
		}

		// send distance vector to all neighbors, they arrive at 0 + Latency(i, j)
		for _, node := range neighbors {
			heap.Push(minHeap, DVMessage{arrivalTime: g.GetWeight(node, i), to: node, from: i, dv: tableEntries})
		}

		// add node to nodeInfo
		nodeInfo = append(nodeInfo, Node{tableEntries})
	}

	return &DVAlgorithm{g, nodeInfo, minHeap}
}

func (d *DVAlgorithm) Run() {
	// while the heap has messages left
	for d.MinHeap.Len() > 0 {
		// pop message off the heap
		currMessage, ok := heap.Pop(d.MinHeap).(DVMessage)
		if !ok {
			panic(fmt.Errorf("Message returned from heap not a DVMessage struct.\n"))
		}
		fmt.Printf("Message from %d received at %d, at time %.2f \n", currMessage.from, currMessage.to, currMessage.arrivalTime)

		// compare current nodes table to the dv entries
		changeMade := false
		for i, tableEntry := range d.NodeInfo[currMessage.to].forwardingTable {
			possibleDistance := currMessage.dv[i].distance + d.g.GetWeight(currMessage.from, currMessage.to)
			if tableEntry.distance > possibleDistance {
				d.NodeInfo[currMessage.to].forwardingTable[i].distance = possibleDistance
				d.NodeInfo[currMessage.to].forwardingTable[i].nextHop = currMessage.from
				changeMade = true
			}
		}

		if changeMade {
			// send out new table to all neighbors
			for _, node := range d.g.GetConnectedNodes(currMessage.to) {
				heap.Push(d.MinHeap, DVMessage{
					arrivalTime: currMessage.arrivalTime + d.g.GetWeight(currMessage.to, node),
					to:          node,
					from:        currMessage.to,
					dv:          d.NodeInfo[currMessage.to].forwardingTable,
				})
			}
		}
	}
}

func (d DVAlgorithm) PrintRoutingTables() {
	for node, nodeData := range d.NodeInfo {
		fmt.Printf("---- Table for node %d ----\n", node)
		for dest, entry := range nodeData.forwardingTable {
			fmt.Printf("Dest: %d; Next Hop: %d; Distance: %.2f; \n", dest, entry.nextHop, entry.distance)
		}
	}
}
