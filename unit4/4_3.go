package main

import "fmt"

func main4_3() {
	var a = [10]int{1, 2, 3, 4, 5}
	reverse1(&a)
	fmt.Println(a)
	b := a[:]
	reverse2(b)
	fmt.Println(a)
}

func reverse1(a *[10]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func reverse2(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
