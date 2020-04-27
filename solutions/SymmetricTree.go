package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func isSymmetric(root *TreeNode) bool {
//
//}

func middleTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	middleTraverse(root.Left)
	fmt.Println(root.Val)
	middleTraverse(root.Right)
}

func main() {
	var root TreeNode
	root = TreeNode{Val: 1}
	root.Left = new(TreeNode)

}
