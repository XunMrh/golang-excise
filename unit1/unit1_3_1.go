package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	file := os.Args[1:]
	if len(file) == 0 {
		countLine(os.Stdin, counts)
	} else {
		for _, arg := range file {
			fp, err := os.Open(arg)
			if err != nil {
				fmt.Printf("err: %s", err)
				continue
			}
			countLine(fp, counts)
		}
	}
	for key, value := range counts {
		fmt.Printf("key: %s, count: %d \n", key, value)
	}
}

func countLine(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
