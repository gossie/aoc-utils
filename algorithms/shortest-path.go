package algorithms

import "container/heap"

type Node struct {
	outgoing, incomming []*edge
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) Add(other *Node, distance int) {
	edge := edge{
		from:   n,
		to:     other,
		length: distance,
	}
	n.outgoing = append(n.outgoing, &edge)
	other.incomming = append(other.incomming, &edge)
}

type edge struct {
	from, to *Node
	length   int
}

type path struct {
	edges []edge
}

func (p *path) TotalDistance() int {
	distance := 0
	for _, e := range p.edges {
		distance += e.length
	}
	return distance
}

type priorityQueue []path

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].TotalDistance() < pq[j].TotalDistance()
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(el any) {
	n := el.(path)
	*pq = append(*pq, n)
}

func (pq *priorityQueue) Pop() any {
	index := pq.Len() - 1
	el := (*pq)[index]
	*pq = (*pq)[0:index]
	return el
}

func ShortestPath(data Node, from, to Node) []Node {
	result := make([]Node, 0)

	pq := make(priorityQueue, 0)
	heap.Init(&pq)

	return result
}
