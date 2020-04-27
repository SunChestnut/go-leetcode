package main

import (
	"fmt"
)

func generate(numRows int) [][]int {

	// 设置特殊值的返回结果
	if numRows == 0 {
		return nil
	}

	var triangle = make([][]int, 1)
	triangle[0] = append(triangle[0], 1)

	for n := 1; n < numRows; n++ {
		var row []int
		var preRow = triangle[n-1]

		// 第一个数字为1
		row = append(row, 1)

		for i := 1; i < n; i++ {
			row = append(row, preRow[i-1]+preRow[i])
		}

		// 最后一个数字为1
		row = append(row, 1)

		triangle = append(triangle, row)
	}
	return triangle
}

func main() {
	res := generate(5)
	fmt.Println(res)
}
