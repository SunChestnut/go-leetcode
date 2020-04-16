package main

// leetcode 237. 删除链表中的节点
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	if node.Next == nil {
		node = nil
		return
	}

	// 将node的下一节点的值赋给node，这样可以直接删除下一节点
	node.Val = node.Next.Val
	delNode := node.Next
	node.Next = delNode.Next
}
