package common

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	var str = ""
	if t == nil {
		return str
	}
	if t.Left != nil {
		str = str + t.Left.String()
		if !strings.HasSuffix(str, ",") {
			str = str + ","
		}
	}
	str = str + strconv.Itoa(t.Val) + ","
	if t.Right != nil {
		str = str + t.Right.String()
		if !strings.HasSuffix(str, ",") {
			str = str + ","
		}
	}
	return str
}

type NodeList []TreeNode

func MakeTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	var depth, count float64
	sv := make([]NodeList, 0)
	var list NodeList
	for _, v := range vals {
		lenth := math.Pow(2, depth)
		if count == 0 {
			list = make([]TreeNode, 0)
		}
		list = append(list, TreeNode{v, nil, nil})
		count = count + 1

		if count == lenth {
			fmt.Printf("depth:%d, list:%v\n", int(depth), list)
			count = 0
			depth = depth + 1
			sv = append(sv, list)
		}
	}
	if count != 0 {
		fmt.Printf("depth:%d, list:%v\n", int(depth), list)
		sv = append(sv, list)
	}

	var cur, next NodeList
	for i := 0; i < len(sv)-1; i++ {
		cur = sv[i]
		next = sv[i+1]
		for j := 0; j < len(cur); j++ {
			if j*2 < len(next) {
				if next[j*2].Val != -1 {
					cur[j].Left = &next[j*2]
				}
			}
			if j*2+1 < len(next) {
				if next[j*2+1].Val != -1 {
					cur[j].Right = &next[j*2+1]
				}
			}
		}
	}
	return &sv[0][0]
}
