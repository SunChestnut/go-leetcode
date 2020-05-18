package main

import "fmt"

// 162. 寻找峰值
func findPeakElement(nums []int) int {

	l := 0
	r := len(nums) - 1

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums))
}
