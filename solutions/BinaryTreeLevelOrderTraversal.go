package main

import "fmt"

type TreeNodeN struct {
	Val   int
	Left  *TreeNodeN
	Right *TreeNodeN
}

// 102. 二叉树的层序遍历
func helper(node *TreeNodeN, level int, res [][]int) [][]int {

	if len(res) == level {
		res = append(res, make([]int, 0))
	}

	res[level] = append(res[level], node.Val)

	if node.Left != nil {
		res = helper(node.Left, level+1, res)
	}

	if node.Right != nil {
		res = helper(node.Right, level+1, res)
	}
	return res
}

func levelOrder(root *TreeNodeN) [][]int {
	var res = make([][]int, 0)
	if root == nil {
		return res
	}
	res = helper(root, 0, res)
	return res
}

func main() {
	var root TreeNodeN
	root.Val = 3
	root.Left = &TreeNodeN{9, nil, nil}
	root.Right = &TreeNodeN{20, &TreeNodeN{15, nil, nil}, &TreeNodeN{7, nil, nil}}

	result1 := levelOrder(&root)

	fmt.Println(result1)

	result2 := levelOrder(nil)
	fmt.Println(result2)
}
