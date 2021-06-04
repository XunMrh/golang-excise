package main

import (
	"fmt"
	"sort"
)

func main7_10() {
	str := str("abccba")
	fmt.Println(isHuiWenChuan(str))
}

type str []byte

func (s str) Len() int {
	return len(s)
}

func (s str) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s str) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func isHuiWenChuan(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
