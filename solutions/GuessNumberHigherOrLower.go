package main

import "fmt"

func guess(num int) int {
	pick := 10
	if pick < num {
		return -1
	} else if pick > num {
		return 1
	} else {
		return 0
	}
}

// 374. 猜数字大小
func guessNumber(n int) int {
	l := 0
	r := n
	mid := -1

	for l <= r {
		mid = l + (r-l)/2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			r = mid - 1
		} else if guess(mid) == 1 {
			l = mid + 1
		}
	}
	return mid
}

func main() {
	fmt.Println(guessNumber(10))
}
