package main

import "fmt"

const (
	_ = (1 << iota)
	kb
	mb
	gb
)

func main3_12() {
	var str1 = "abc"
	var str2 = "hanson"
	table := make(map[rune]int)
	for _, v := range str1 {
		table[v]++
	}
	for _, v := range str2 {
		if table[v] == 0 {
			fmt.Println("false")
			return
		}
	}
	fmt.Println("true")
}
