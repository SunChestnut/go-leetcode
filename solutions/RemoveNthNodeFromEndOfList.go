package main

import "fmt"

// leetcode 19. 删除链表的倒数第N个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// 设置虚拟头节点 dummyHead
	var dummyHead ListNode
	dummyHead.Val = 0
	dummyHead.Next = head

	var p = &dummyHead
	var q = &dummyHead

	// 使得p和q中间间隔n个点
	for i := 0; i <= n; i++ {
		q = q.Next
	}

	// 移动指针p和q，使q指向nil,q指向待删除节点的上一个节点
	for q != nil {
		p = p.Next
		q = q.Next
	}

	var delNode = p.Next
	p.Next = delNode.Next

	return dummyHead.Next
}

func main() {
	var head ListNode
	head.Val = 1

	var next1 ListNode
	next1.Val = 2
	head.Next = &next1

	var next2 ListNode
	next2.Val = 3
	next1.Next = &next2

	var next3 ListNode
	next3.Val = 4
	next2.Next = &next3

	var next4 ListNode
	next4.Val = 5
	next3.Next = &next4

	end := removeNthFromEnd(&head, 2)

	for end != nil {
		fmt.Println(end.Val)
		end = end.Next
	}
}
