package main

import (
	"crypto/sha256"
	"fmt"
)

var table [256]byte

func getTable() {
	for i := 1; i < len(table); i++ {
		table[i] = table[i/2] + byte(i&1)
		fmt.Println(table[i])
	}
}

func main4_1() {
	getTable()
	c1 := sha256.Sum256([]byte{'x'})
	var count = 0
	for _, v := range c1 {
		count += int(table[v])
	}
	fmt.Printf("count: %d\n", count)
}
