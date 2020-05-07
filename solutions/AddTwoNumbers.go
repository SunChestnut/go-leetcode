package main

import (
	"fmt"
	"go-leetcode/entity"
)

// leetcode 02. 两数相加
func addTwoNumbers(l1 *entity.ListNode, l2 *entity.ListNode) *entity.ListNode {

	// 创建新的ListNode存储结果
	res := new(entity.ListNode)
	// 创建cur指针指向结果链表
	var cur = res

	// 创建指针p和q分别指向l1和l2
	var p = l1
	var q = l2

	carry := 0
	// 2->4->3 + 5->6->4
	for p != nil || q != nil {
		v1 := 0
		v2 := 0
		if p != nil {
			v1 = p.Val
		}
		if q != nil {
			v2 = q.Val
		}

		sum := v1 + v2 + carry
		carry = sum / 10

		cur.Next = &entity.ListNode{Val: sum % 10}
		cur = cur.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 {
		cur.Next = &entity.ListNode{Val: carry, Next: nil}
	}
	return res.Next
}

func main() {
	l1 := entity.ListNode{Val: 5, Next: nil}
	l2 := entity.ListNode{Val: 5, Next: nil}

	res := addTwoNumbers(&l1, &l2)
	fmt.Printf("%v", res)
}
