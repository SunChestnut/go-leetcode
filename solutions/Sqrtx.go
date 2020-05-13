package main

import (
	"fmt"
	"time"
)

// 69. x 的平方根
func mySqrt(x int) int {
	l := 0
	r := x
	ans := -1

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

func main() {
	startTime := time.Now()
	fmt.Println(mySqrt(2147395599))
	finishTime := time.Now()

	operate := finishTime.Sub(startTime)
	fmt.Println(operate)
}
