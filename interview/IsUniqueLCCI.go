package main

import (
	"fmt"
)

// 面试题 01.01. 判定字符是否唯一
func isUnique(astr string) bool {

	var arr [26]int

	for _, v := range astr {
		index := v - 'a'
		if arr[index] == 0 {
			arr[index] = 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isUnique("abc"))
	fmt.Println(isUnique("leetcode"))
}
