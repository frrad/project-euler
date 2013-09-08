package main

import (
	"euler"
	"fmt"
	"strconv"
	"strings"
)

func getNum(a string) int {
	probLen := 8 //Length of `Problem '

	starts := strings.Index(a, "Problem ")
	ends := strings.Index(a[starts+probLen:], " ")

	isolated := a[starts+probLen : starts+probLen+ends]
	number, _ := strconv.Atoi(isolated)

	return number
}

func luckySeive(max int) []int {

	luckyseive := make([]int, max)
	for i := 0; i < max; i++ {
		luckyseive[i] = i + 1
	}

	last := -1
	pointer := 1

	for pointer < len(luckyseive) {

		last = luckyseive[pointer]

		for del := last - 1; del < len(luckyseive); del += last {
			luckyseive[del] = 0
		}
		for i := 0; i < len(luckyseive); i++ {
			if luckyseive[i] == 0 {
				luckyseive = append(luckyseive[:i], luckyseive[i+1:]...)
				i--
			}

		}

		if luckyseive[pointer] == last {
			pointer++
		}

	}

	return luckyseive
}

func main() {
	lineL := 60

	page := euler.Import("../problemdata/Project Euler.html")

	max := -1 //number of problems total
	dict := make(map[int]bool)

	for _, line := range page {
		split := strings.Split(line, "<td style=\"width:20px;height:20px;vertical-align:middle;background-color:#")
		for _, prob := range split {

			if len(prob) > 5 {
				if prob[:6] == "CEE7B6" {
					//Green = Complete

					number := getNum(prob)
					dict[number] = true
					if number > max {
						max = number
					}

				} else if prob[:3] == "fff" {
					//White = Incomplete

					number := getNum(prob)
					if number > max {
						max = number
					}
				}
			}

		}
	}

	/*
		//Set all to complete for testing
		for i := 1; i <= max; i++ {
			dict[i] = true
		}
	*/

	done := 0
	prime := 0
	triangle := 0
	lucky := 0
	dec2 := 0

	for i := 1; i <= 25; i++ {
		if dict[i*(i+1)/2] {
			triangle++
		}
	}

	for i := 1; i <= max; i++ {
		if dict[i] {
			done++
			if euler.IsPrime(int64(i)) {
				prime++
			}
		}
	}

	luckyseive := luckySeive(max)
	for i := 0; i < len(luckyseive); i++ {
		if dict[luckyseive[i]] {
			lucky++
		}
	}

	for i := 0; i < 10; i++ {
		here := 0
		for j := 200 + 10*i + 1; j < 200+10*(i+1)+1; j++ {
			if dict[j] {
				here++
			}
		}

		if here > 0 {
			dec2++
		}
	}

	fmt.Println("Done", done, "/", max, " problems")

	for i := 0; i < lineL; i++ {
		fmt.Print("=")
	}
	fmt.Print("\n")

	for i := 1; i <= max; i++ {
		if dict[i] {
			fmt.Print("X")
		} else {
			fmt.Print(" ")
		}

		if i%lineL == 0 {
			fmt.Print("\n")
		}
	}
	if max%lineL != 0 {
		fmt.Print("\n")
	}
	for i := 0; i < lineL; i++ {
		fmt.Print("=")
	}
	fmt.Print("\n\n")

	fmt.Println("Prime Obsession:\t", prime, "/ 50 prime numbered problems")
	fmt.Println("Triangle Trophy:\t", triangle, "/ 25 first triangle numbered problems")
	fmt.Println("Lucky Luke:\t\t", lucky, "/ 50 lucky numbered problems")
	fmt.Println("Decimation II:\t\t", dec2, "/ 10 rows")

	fmt.Print("\n")

	track := make(map[int]int)

	if prime < 50 {
		fmt.Print("Primes: ")
		for i := 1; i < max; i++ {
			if !dict[i] && euler.IsPrime(int64(i)) {
				fmt.Print(i, " ")
				track[i]++
			}
		}
		fmt.Print("\n")
	}

	if triangle < 25 {
		fmt.Print("Triangle Numbers: ")

		for i := 1; i <= 25; i++ {
			if !dict[i*(i+1)/2] {
				fmt.Print(i*(i+1)/2, " ")
				track[i*(i+1)/2]++
			}
		}

		fmt.Print("\n")

	}

	if lucky < 50 {
		fmt.Print("Lucky Numbers: ")

		for i := 0; i < len(luckyseive); i++ {
			if !dict[luckyseive[i]] {
				fmt.Print(luckyseive[i], " ")
				track[luckyseive[i]]++
			}
		}

		fmt.Print("\n")
	}

	if dec2 < 10 {
		fmt.Print("Decimation II: ")

		for i := 0; i < 10; i++ {
			here := 0
			for j := 200 + 10*i + 1; j < 200+10*(i+1)+1; j++ {
				if dict[j] {
					here++
				}
			}

			if here == 0 {
				for j := 200 + 10*i + 1; j < 200+10*(i+1)+1; j++ {
					fmt.Print(j, " ")
					track[j]++
				}
			}
		}

		fmt.Print("\n")
	}

	maxTrack := -1

	fmt.Print("\nRepeats: ")

	for i := 1; i <= max; i++ {
		if track[i] > 1 {
			fmt.Print(i, " ")

		}
		if track[i] > maxTrack {
			maxTrack = track[i]
		}
	}

	fmt.Print("\n")

	fmt.Print("Most Repeated: ")

	for i := 1; i <= max; i++ {
		if track[i] == maxTrack {
			fmt.Print(i, " ")

		}

	}

	fmt.Print("\n")

}
