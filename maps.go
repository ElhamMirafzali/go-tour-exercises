package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, w := range words {
		m[w] += 1
	}
	return m
}

func main() {
	fmt.Println(WordCount("I am learning Go!"))
	fmt.Println(WordCount("Hi You Hi Again"))
}
