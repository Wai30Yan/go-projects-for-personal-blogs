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
	findNode(node, "class")
	parseIndent(node, startElement, endElement)
}

func findNode(n *html.Node, attrKey string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == attrKey {
				fmt.Println(attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findNode(c, attrKey)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func parseIndent(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseIndent(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}