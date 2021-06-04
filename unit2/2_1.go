package main

import (
	"fmt"

	"example.com/m/tempconv"
)

func main2_1() {
	fmt.Println(tempconv.AbsoluteC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.CToK(tempconv.FreezingC))
}
