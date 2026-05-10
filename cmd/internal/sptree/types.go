package sptree

// Message structs store a Spanning Tree Protocol (STP) message incuding the proposed LAN root, the distance from X to the root, and the sender of the message X.
type Message struct {
	root     int
	distance int
	from     int
}

// NodeData structs store globally sent and recieved messages, the next hop, current root, and current distance to the root, for each node in the LAN.
type NodeData struct {
	sent         []Message
	recieved     []Message
	nextHop      int
	currRoot     int
	currDistance int
}
