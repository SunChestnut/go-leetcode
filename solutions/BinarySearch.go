package main

import "fmt"

// 二分查找
func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	mid := len(nums) / 2

	for left <= right {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	res := search(nums, 13)
	fmt.Println(res)
}
