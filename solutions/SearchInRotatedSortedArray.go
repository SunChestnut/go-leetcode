package main

import "fmt"

// leetcode33. 搜索旋转排序数组
func searchInRotatedSortedArray(nums []int, target int) int {
	n := len(nums)

	if n == 0 {
		return -1
	}

	if n == 1 && nums[0] == target {
		return 0
	}

	l := 0
	r := n - 1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		// 判断 nums[0...mid]是否是递增的
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{5, 1, 3}
	fmt.Println(searchInRotatedSortedArray(nums, 3))
}
