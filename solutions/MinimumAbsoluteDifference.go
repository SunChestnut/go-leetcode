package main

import (
	"fmt"
	"math"
	"sort"
)

// leetcode 1200. 最小绝对差
func minimumAbsDifference(arr []int) [][]int {

	var n = len(arr)

	if n == 0 {
		return nil
	}

	sort.Ints(arr)
	var minSub = math.MaxInt32
	for i := 1; i < n; i++ {
		minSub = int(math.Min(float64(minSub), math.Abs(float64(arr[i]-arr[i-1]))))
	}

	var res [][]int
	for j := 1; j < n; j++ {
		if int(math.Abs(float64(arr[j]-arr[j-1]))) == minSub {
			tempArr := []int{arr[j-1], arr[j]}
			res = append(res, tempArr)
		}
	}
	return res
}

func main() {
	arr := []int{3, 8, -10, 23, 19, -4, -14, 27}
	difference := minimumAbsDifference(arr)
	fmt.Println(difference)
}
