package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main3_10() {
	var buf bytes.Buffer
	buf.WriteByte('[')
	inputScan := bufio.NewScanner(os.Stdin)
	inputScan.Scan()
	line := inputScan.Text()
	for i := 0; i < len(line); {
		fmt.Println(i)
		if (i + 3) > len(line) {
			buf.WriteString(line[i:])
			break
		} else {
			buf.WriteString(line[i : i+3])
			buf.WriteByte(',')
			i += 3
		}
	}
	buf.WriteByte(']')
	fmt.Println(buf.String())
}
