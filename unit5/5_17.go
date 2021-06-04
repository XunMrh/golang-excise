package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main5_17() {
	url := "http://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	nodes := getSameName(doc, "img")
	for _, node := range nodes {
		fmt.Printf("data: %v", node.Data)
		fmt.Printf("attri: %v\n", node.Attr)
	}
}

func getSameName(d *html.Node, names ...string) []*html.Node {
	var nodes []*html.Node
	if d.Type == html.ElementNode {
		for _, name := range names {
			if d.Data == name {
				nodes = append(nodes, d)
				break
			}
		}
	}
	for node := d.FirstChild; node != nil; node = node.NextSibling {
		nodes = append(nodes, getSameName(node, names...)...)
	}
	return nodes
}
