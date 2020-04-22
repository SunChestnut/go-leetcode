package main

import "fmt"

// leetcode 168. Excel表列名称
func convertToTitle(n int) string {
	res := ""

	for n > 0 {
		n -= 1
		res = string('A'+(n%26)) + res
		n /= 26
	}

	return res
}

func main() {
	fmt.Println(convertToTitle(28))
}
