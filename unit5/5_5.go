package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main5_5() {
	word, img, err := CountNum("https://www.sulinehk.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
	}
	fmt.Printf("word num: %d, image num: %d \n", word, img)
}

func CountNum(url string) (word, img int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	return word, img, countNum(doc, &word, &img)
}

func countNum(n *html.Node, word, img *int) (err error) {
	if n.Type == html.TextNode {
		*word += len(strings.Fields(n.Data))
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		*img++
	}
	for node := n.FirstChild; node != nil; node = node.NextSibling {
		countNum(node, word, img)
	}
	return nil
}
