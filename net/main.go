package main

import (
	"io"
	"log"
	"net/http"
	"os"

	// "golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://wai30yan.github.io/personal-blog/blogs/nmap/")
	if err != nil {
		log.Panic(err)
	}

	b, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	writeFile(b)



}

func readFile(fileName string) {

}

func writeFile(b []byte) {
	f, _ := os.Open("response.txt")
	f.Write(b)
	f.Close()
}