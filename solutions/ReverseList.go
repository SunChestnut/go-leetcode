package main

import "fmt"

// leetcode 206. 反转链表
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var cuNode = reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return cuNode
}

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	var head ListNode
	head.Val = 0
	head.Next = nil

	var curNode = &head
	for v, _ := range arr {
		fmt.Println(v, curNode)
		curNode = curNode.Next
		curNode.Val = v
	}

	fmt.Println("curNode = ", curNode)

	resNode := reverseList(head.Next)
	var cur2 = &resNode
	for cur2 != nil {
		fmt.Println((*cur2).Val)
		cur2 = &(*cur2).Next
	}
}
