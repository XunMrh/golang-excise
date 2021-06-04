package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main4_2() {
	var f = flag.String("sha", "sha256", "-sha256, -sha384, -sha512")
	flag.Parse()
	switch *f {
	case "sha256":
		fmt.Printf("sha256 of x: %x\n", sha256.Sum256([]byte{'x'}))
	case "sha384":
		fmt.Printf("sha384 of x: %x\n", sha512.Sum384([]byte{'x'}))
	case "sha512":
		fmt.Printf("sha512 of x: %x\n", sha512.Sum512([]byte{'x'}))
	}
}
