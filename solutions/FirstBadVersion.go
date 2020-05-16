package main

import "fmt"

func isBadVersion(version int) bool {
	return version != 1 && version != 2 && version != 3
}

// 278. 第一个错误的版本
// 二分法
func firstBadVersion(n int) int {

	l := 1
	r := n

	for l < r {
		mid := l + (r-l)/2
		// isBadVersion(mid) == true 即为坏版本
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(firstBadVersion(5))
}
