package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main5_2() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
	}
	for k, v := range VisitAll(doc) {
		fmt.Printf("node: %s, number: %d", k, v)
	}
}

var count = make(map[string]int)

func VisitAll(n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	for node := n.FirstChild; node != nil; node = node.NextSibling {
		VisitAll(node)
	}
	return count
}
