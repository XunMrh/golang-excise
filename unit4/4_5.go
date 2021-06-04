package main

import "fmt"

func main4_5() {
	var str = []string{"a", "b", "b", "c", "c"}
	mslice := str[:]
	mslice = unique(mslice)
	fmt.Println(mslice)
}

func unique(str []string) []string {
	var count = 0
	for i := 1; i < len(str); i++ {
		if str[i] != str[count] {
			count++
			str[count] = str[i]
		}
	}
	return str[:count+1]
}
