package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func main7_2() {
	w, c := countWriter(ioutil.Discard)
	fmt.Fprintf(w, "hello, world\n")
	fmt.Println(c)
}

type newWriter struct {
	w io.Writer
	c int
}

func (nw *newWriter) Write(p []byte) (int, error) {
	n, err := nw.w.Write(p)
	nw.c += n
	return nw.c, err
}

func countWriter(w io.Writer) (io.Writer, int) {
	nw := newWriter{w, 0}
	return &nw, nw.c
}
