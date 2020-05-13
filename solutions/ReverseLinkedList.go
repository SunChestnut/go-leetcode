package main

import (
	"fmt"
	"go-leetcode/entity"
)

func reverseListII(head *entity.ListNode) *entity.ListNode {
	var newHead *entity.ListNode = nil
	curHead := head

	for curHead != nil {
		next := curHead.Next
		curHead.Next = newHead

		newHead = curHead
		curHead = next
	}
	return newHead
}

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	head := entity.ListNode{}
	var curHead = &head
	for _, v := range array {
		curHead.Next = &entity.ListNode{Val: v}
		curHead = curHead.Next
	}

	newHead := reverseListII(head.Next)
	fmt.Printf("%+v", newHead)
}
