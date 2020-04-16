package main

import (
	"fmt"
	"math"
)

func dominantIndex(nums []int) int {

	var bigNum = math.MinInt32
	var finalIndex = 0
	for i, v := range nums {
		if bigNum < v {
			bigNum = v
			finalIndex = i
		}
	}

	for _, v := range nums {
		if bigNum != v && bigNum < v*2 {
			return -1
		}
	}

	return finalIndex
}

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
}
