package main

import (
	"fmt"
	"strings"
)

func Jion(args ...string) string {
	if len(args) < 2 {
		return ""
	}
	space := args[len(args)-1]
	last := args[len(args)-2]
	var b strings.Builder
	tempString := args[:len(args)-2]
	for _, str := range tempString {
		b.WriteString(str)
		b.WriteString(space)
	}
	b.WriteString(last)
	return b.String()
}

func main5_16() {
	fmt.Println(Jion("aa", "bb", "|"))
}
