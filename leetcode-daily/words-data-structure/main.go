package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"flower", "flower", "flower", "flower"}
	strs1 := []string{"dog", "racecar", "car"}

	w := WordDictionary{strs}
	w2 := WordDictionary{strs1}

	fmt.Println(w.Search(".lower"))
	fmt.Println(w2.Search("..cecar"))
}

type WordDictionary struct {
	words []string
}

func Constructor() WordDictionary {
	return WordDictionary{
		words: make([]string, 0),
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.words = append(this.words, word)
}

func (this *WordDictionary) Search(word string) bool {
	result := ""
	
	if strings.Contains(word, ".") {
		fmt.Println(strings.Count(word, "."))
	}
	for _, l := range word {
		result += string(l)
		for _, w := range this.words {
			if strings.HasPrefix(w, result) && result == w {
				return true					
			} 
		}
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */