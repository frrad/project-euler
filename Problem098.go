package main

import (
	"./euler"
	"fmt"
	"strings"
	"time"
)

func main() {
	starttime := time.Now()

	data := euler.Import("problemdata/words.txt")
	line := data[0]
	words := strings.Split(line, ",")

	anagrams := make(map[string][]string)

	for _, word := range words {
		stripped := word[1 : len(word)-1]
		sorted := euler.BubbleSort(stripped)

		anagrams[sorted] = append(anagrams[sorted], stripped)
	}

	for _, set := range anagrams {

		if len(set) > 1 {
			fmt.Println(set)
		}
	}

	fmt.Println("Elapsed time:", time.Since(starttime))

}
