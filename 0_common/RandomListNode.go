package common

// RandomListNode List Node with Random Pointer
type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

// CreateRandomList Create random list node with given values
func CreateRandomList(vals [][]int) *RandomListNode {
	nodes := make([]*RandomListNode, len(vals))
	for i := range vals {
		nodes[i] = &RandomListNode{vals[i][0], nil, nil}
	}
	for i := range vals {
		if i < len(vals)-1 {
			nodes[i].Next = nodes[i+1]
		}
		if vals[i][1] >= 0 {
			nodes[i].Random = nodes[vals[i][1]]
		}
	}
	return nodes[0]
}
