package main

// leetcode 234. 回文链表
func isPalindrome(head *ListNode) bool {

	if head == nil {
		return true
	}

	var n = 0
	var cur1 = head
	for cur1 != nil {
		n++
		cur1 = cur1.Next
	}

	var arr = make([]int, n)
	var cur2 = head
	var i = 0
	for cur2 != nil {
		arr[i] = cur2.Val
		i++
		cur2 = cur2.Next
	}

	var r = n - 1
	for l := 0; l < n/2; l++ {
		if arr[l] == arr[r] {
			r--
		} else {
			return false
		}
	}
	return true
}

func main() {
	isPalindrome(nil)
}
