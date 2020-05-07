package main

import (
	common "LeetCodeByGo/0_common"
	"fmt"
	"strconv"
	"strings"
)

// TreeNode Node of Binary Tree
type TreeNode = common.TreeNode

func main() {
	// data := "[1,2,3,null,null,4,5]"
	data := "[5,2,3,null,null,2,4,3,1]"
	codec := Constructor()
	root := codec.deserialize(data)
	fmt.Println("Deserialize Output:", root)
	fmt.Println("Serialize Output:", codec.serialize(root))
}

// Codec serialize and deserialize a binary tree
type Codec struct {
}

// Constructor create codec
func Constructor() Codec {
	return Codec{}
}

// serialize Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	res := ""
	nodes := []*TreeNode{root}
	buf := []*TreeNode{}
	for {
		str := ""
		for i := range nodes {
			if nodes[i] == nil {
				str += "null,"
			} else {
				str += strconv.Itoa(nodes[i].Val) + ","
				buf = append(buf, []*TreeNode{nodes[i].Left, nodes[i].Right}...)
			}
		}
		if len(buf) == 0 {
			break
		}
		res += str
		nodes = make([]*TreeNode, len(buf))
		copy(nodes, buf)
		buf = make([]*TreeNode, 0)
	}
	for len(res) > 5 && res[len(res)-5:] == "null," {
		res = res[0 : len(res)-5]
	}
	res = "[" + res[0:len(res)-1] + "]"
	return res
}

// deserialize Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 2 {
		return nil
	}
	data = strings.TrimSpace(data[1 : len(data)-1])
	vals := strings.Split(data, ",")
	if len(vals) == 0 {
		return nil
	}
	v, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: v, Left: nil, Right: nil}
	nodes := []*TreeNode{root}
	buf := []*TreeNode{}
	i := 1
	var node *TreeNode
	for i < len(vals) && len(nodes) > 0 {
		for j := range nodes {
			if nodes[j] == nil {
				continue
			}
			//Create Left Leaf
			node = createNode(vals[i])
			buf = append(buf, node)
			nodes[j].Left = node

			i++
			if i == len(vals) {
				break
			}

			//Create Right Leaf
			node = createNode(vals[i])
			buf = append(buf, node)
			nodes[j].Right = node

			i++
			if i == len(vals) {
				break
			}
		}

		nodes = make([]*TreeNode, len(buf))
		copy(nodes, buf)
		buf = make([]*TreeNode, 0)
	}
	return root
}

func createNode(val string) *TreeNode {
	if val == "" || val == "null" {
		return nil
	}
	v, _ := strconv.Atoi(val)
	return &TreeNode{Val: v, Left: nil, Right: nil}
}
