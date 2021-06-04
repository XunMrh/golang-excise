package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main4_6() {
	var ms = []byte("我是  中 国 人")
	ms = rmSpace(ms)
	fmt.Printf("%s\n", ms)
}

func rmSpace(ms []byte) []byte {
	for i := 0; i < len(ms); {
		first, size := utf8.DecodeRune(ms[i:])
		if unicode.IsSpace(first) {
			second, size2 := utf8.DecodeRune(ms[i+size:])
			if unicode.IsSpace(second) {
				ms[i] = ' '
				copy(ms[i+1:], ms[i+size+size2:])
				ms = ms[:len(ms)-size-size2+1]
			}
		}
		i++
	}
	return ms
}
