package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main7_1() {
	str := []byte("hello, world")
	var wc wordCounter = 0
	var lc lineCounter = 0
	fmt.Fprintln(&wc, str)
	fmt.Println(wc)
	fmt.Fprintln(&lc, str)
	fmt.Println(lc)
}

type wordCounter int

func (wc *wordCounter) Write(p []byte) (int, error) {
	mScanner := bufio.NewScanner(bytes.NewReader(p))
	mScanner.Split(bufio.ScanWords)
	for mScanner.Scan() {
		*wc++
	}
	return len(p), nil
}

type lineCounter int

func (lc *lineCounter) Write(p []byte) (int, error) {
	mScanner := bufio.NewScanner(bytes.NewReader(p))
	mScanner.Split(bufio.ScanLines)
	for mScanner.Scan() {
		*lc++
	}
	return len(p), nil
}
