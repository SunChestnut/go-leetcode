package main

// leetcode 21. 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	// 维护一个哨兵节点
	var preHead ListNode
	preHead.Val = -1
	preHead.Next = nil

	var pre = &preHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}

	pre.Next = l1
	if pre.Next == nil {
		pre.Next = l2
	}
	return preHead.Next
}
