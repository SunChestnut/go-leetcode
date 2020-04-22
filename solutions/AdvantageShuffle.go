package main

import (
	"fmt"
	"sort"
)

// leetcode 870. 优势洗牌
func advantageCount(A []int, B []int) []int {
	var sortedA = make([]int, len(A))
	copy(sortedA, A)
	sort.Ints(sortedA)

	var sortedB = make([]int, len(B))
	copy(sortedB, B)
	sort.Ints(sortedB)

	assigned := make(map[int][]int)
	for _, v := range sortedB {
		slice := make([]int, 0)
		assigned[v] = slice
	}

	var remaining = make([]int, 0)

	var j = 0
	for _, a := range sortedA {
		if a > sortedB[j] {
			assigned[sortedB[j]] = append(assigned[sortedB[j]], a)
			j++
		} else {
			remaining = append(remaining, a)
		}
	}

	var res = make([]int, len(B))
	for i, v := range B {
		if len(assigned[v]) > 0 {
			res[i] = assigned[v][0]
			assigned[v] = assigned[v][1:]
		} else {
			res[i] = remaining[0]
			remaining = remaining[1:]
		}
	}
	return res
}

func main() {
	A := []int{12, 24, 8, 32}
	B := []int{13, 25, 32, 11}

	res := advantageCount(A, B)
	fmt.Println(res)
}
