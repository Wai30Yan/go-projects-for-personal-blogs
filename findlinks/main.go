package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, _ := http.Get("https://wai30yan.github.io/personal-blog/tags/")
	// resp, _ := http.Get("https://arc.dev/developer-blog/web-developer-portfolio/")

	fmt.Println(resp.Request.URL)
	fmt.Println(resp.Request.URL.Path)
	fmt.Println(resp.Request.URL.Host)
	fmt.Println("-------------------------")

	node, _ := html.Parse(resp.Body)
	var links []string
	var findLinks func(n *html.Node)
	findLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link, _ := resp.Request.URL.Parse(attr.Val)
					fmt.Println(attr.Val)
					links = append(links, link.String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findLinks(c)
		}
	}

	findLinks(node)

	fmt.Println("-------------------------")
	fmt.Println(len(links))
}