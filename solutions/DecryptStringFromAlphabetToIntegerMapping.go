package main

import "fmt"

// 1309. 解码字母到整数映射
func freqAlphabets(s string) string {

	res := ""
	// pos := 'a' - 1

	runes := []rune(s)
	for i, v := range runes {
		fmt.Println(i, string(v))
	}

	return res
}

func main() {
	//s := '3'
	//pos := 'a' - 1
	//target := s - '0' + pos
	//fmt.Printf("%s", string(target))
	//fmt.Println()
	//fmt.Printf("v = %s", string('3'-'0'+pos))
	//fmt.Println()

	fmt.Printf("%s", freqAlphabets("15#325#"))
}
