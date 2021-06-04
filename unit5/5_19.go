package main

import "fmt"

func main5_19() {
	fmt.Println(getNoZero())
}

func getNoZero() (result int) {
	defer func() {
		result = 3
		_ = recover()
		fmt.Println("after recover ")
	}()
	fmt.Println("before panic")
	panic("panic")
}
