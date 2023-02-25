package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

var url = "https://wai30yan.github.io/personal-blog/"

func main() {
	// f, _ := os.Open("index.html")
	resp, _ := http.Get(url)
	node, _ := html.Parse(resp.Body)
	findNode(node)
}

func findNode(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "class" {
				fmt.Println(attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findNode(c)
	}
}