package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := "https://wai30yan.github.io/personal-blog/"

	resp, _ := http.Get(url)

	b, _ := io.ReadAll(resp.Body)

	f, _ := os.Create("index.txt")

	f.Write(b)

	node, _ := html.Parse(resp.Body)
	c, _ := os.Create("copy.txt")
	io.Copy(c, f)
	fmt.Println(node)
}