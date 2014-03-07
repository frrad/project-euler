package main

import (
	"euler"
	"fmt"
	"sort"
	"strings"
	"time"
)

func score(name string) (total int) {
	for _, let := range name {
		total += int(let - 64)
	}
	return total
}

func main() {
	starttime := time.Now()

	data := euler.Import("../problemdata/names.txt")
	names := strings.Split(data[0], ",")
	for i, name := range names {
		names[i] = strings.Trim(name, "\"")
	}

	sort.Strings(names)

	total := 0
	for i, name := range names {
		total += (i + 1) * score(name)
	}

	fmt.Println(total)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
