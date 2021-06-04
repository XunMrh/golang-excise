package main

import (
	"fmt"
	//"os"
)

func renameArg0(args []string) {
	args[0] = "newname"
	fmt.Println(args[0])
}
