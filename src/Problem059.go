package main

import (
	"euler"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var ciphertext []int

func decode(password []int) (cleartext []int) {
	cleartext = make([]int, len(ciphertext))
	for i, letter := range ciphertext {
		cleartext[i] = letter ^ password[i%len(password)]
	}
	return
}

func main() {
	starttime := time.Now()

	ciphertext = make([]int, 0)

	in := euler.Import("../problemdata/cipher1.txt")
	split := strings.Split(in[0], ",")

	for _, number := range split {
		integer, _ := strconv.Atoi(number)
		ciphertext = append(ciphertext, integer)
	}

	max := 0
	key := make([]int, 3)

	for i := 97; i <= 122; i++ {
		for j := 97; j <= 122; j++ {

			for k := 97; k <= 122; k++ {

				try := []int{i, j, k}

				cleartext := decode(try)

				score := 0

				for _, letter := range cleartext {

					if letter == 32 {
						score++
					}
				}

				if score > max {
					max = score
					key = []int{i, j, k}
				}
			}
		}
	}

	fmt.Print("Key: ", key, " = ")
	for _, letter := range key {
		fmt.Print(string(letter))
	}
	fmt.Print("\n")

	cleartext := decode(key)

	sum := 0
	for _, letter := range cleartext {
		fmt.Print(string(letter))
		sum += letter
	}

	fmt.Printf("\n%d\n", sum)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
