package sptree

import (
	"fmt"

	"github.com/eboyden42/distributed-algorithms/cmd/internal/graph"
)

// SPTreeAlgorithm structs store a graph to get connections from, and NodeData (as defined in graph/types) for each node.
type SPTreeAlgorithm struct {
	g         graph.Graph[float64]
	roundData []NodeData
}

func New(g graph.Graph[float64]) SPTreeAlgorithm {
	n := g.Len()
	roundData := []NodeData{}
	for i := range n {
		nodeData := NodeData{
			[]Message{{i, 0, i}},
			[]Message{},
			i,
			i,
			0,
		}
		roundData = append(roundData, nodeData)
	}
	return SPTreeAlgorithm{g, roundData}
}

func (s *SPTreeAlgorithm) recieve() {
	// iterate through each node
	for i := range len(s.roundData) {
		// clear recieved slice
		s.roundData[i].recieved = nil
	}
	// iterate through each node
	for i, nodeData := range s.roundData {
		// for each message in sent
		for _, message := range nodeData.sent {
			// move it to recieved for connected nodes (based on the graph)
			for _, node := range s.g.GetConnectedNodes(i) {
				s.roundData[node].recieved = append(s.roundData[node].recieved, message)
			}
		}
	}
}

func (s *SPTreeAlgorithm) send() {
	for i := range len(s.roundData) {
		// clear sent slice
		s.roundData[i].sent = nil
	}
	// iterate through each node
	for i := range len(s.roundData) {
		// for each message in recieved
		for _, message := range s.roundData[i].recieved {
			// if message has a root node lower than yours, or message has the same root but a shorter distance
			if message.root < s.roundData[i].currRoot || (message.root == s.roundData[i].currRoot && message.distance+1 < s.roundData[i].currDistance) {
				// set next hop to from, set root to proposed root
				s.roundData[i].nextHop = message.from
				s.roundData[i].currRoot = message.root
				s.roundData[i].currDistance = message.distance + 1
				// add one distance, and add new message to send
				s.roundData[i].sent = append(s.roundData[i].sent, Message{message.root, message.distance + 1, i})
			}
		}
	}
}

func (s SPTreeAlgorithm) PrintRoundData(roundNumber int) {
	fmt.Printf("-------------(Round %d)--------------\n", roundNumber)
	for i, nodeData := range s.roundData {
		fmt.Printf("%d: Sent: %v; Recvd: %v; Next Hop: %d; Root: %d \n", i, nodeData.sent, nodeData.recieved, nodeData.nextHop, nodeData.currRoot)
	}
}

func (s SPTreeAlgorithm) messagesAreLeft() bool {
	for _, nodeData := range s.roundData {
		if len(nodeData.recieved) != 0 || len(nodeData.sent) != 0 {
			return true
		}
	}
	return false
}

func (s *SPTreeAlgorithm) Run() {
	s.PrintRoundData(0)
	for i := 1; s.messagesAreLeft(); i++ {
		s.recieve()
		s.send()
		s.PrintRoundData(i)
	}
}
