package main

import (
	"fmt"
	"os"
)

func printArgs(args []string) {
	for index, arg := range args {
		fmt.Println(index, arg)
	}
}

func maint() {
	printArgs(os.Args)
	renameArg0(os.Args)
}
