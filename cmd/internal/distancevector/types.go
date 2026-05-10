package distancevector

type DVMessage struct {
	arrivalTime float64
	to          int
	from        int
	dv          []TableEntry
}

func (m DVMessage) Evaluate() float64 {
	return m.arrivalTime
}

type TableEntry struct {
	nextHop  int
	distance float64
}

type Node struct {
	forwardingTable []TableEntry
}
