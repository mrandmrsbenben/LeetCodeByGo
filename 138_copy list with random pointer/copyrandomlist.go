package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
)

// Node Random List Node
type Node = common.RandomListNode

func main() {
	vals := [][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	head := copyRandomList(common.CreateRandomList(vals))
	for head != nil {
		fmt.Println("Output: ", head)
		head = head.Next
	}

}

//执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗 :3.6 MB, 在所有 Go 提交中击败了100.00%的用户
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodes, newNodes := make([]*Node, 0), make([]*Node, 0)
	nm := make(map[string]int)
	cnt := 0
	for head != nil {
		nodes = append(nodes, head)
		nm[fmt.Sprintf("%p", nodes[cnt])] = len(nm)
		newNodes = append(newNodes, &Node{Val: head.Val, Next: nil, Random: nil})
		head = head.Next
		cnt++
	}

	var index int
	for i := range newNodes {
		if i < len(newNodes)-1 {
			newNodes[i].Next = newNodes[i+1]
		}
		if nodes[i].Random != nil {
			index = nm[fmt.Sprintf("%p", nodes[i].Random)]
			newNodes[i].Random = newNodes[index]
		}
	}
	return newNodes[0]
}
