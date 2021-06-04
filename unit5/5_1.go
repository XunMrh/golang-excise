package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main5_1() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(link []string, n *html.Node) []string {
	if n == nil {
		return link
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				link = append(link, a.Val)
			}
		}
	}
	link = visit(link, n.FirstChild)
	link = visit(link, n.NextSibling)
	return link
}
