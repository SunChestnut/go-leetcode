package main

import (
	"fmt"
	"strings"
)

// LeetCode 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, v := range strs {
		for strings.Index(v, prefix) != 0 {
			prefix = prefix[0 : len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Println(result)
}
