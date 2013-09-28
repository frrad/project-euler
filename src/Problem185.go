package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// const (
// 	guessN = 6
// 	height = 5
// )

// var (
// 	gS = []string{"90342", "70794", "39458", "34109", "51545", "12531"}
// 	gI = []int{2, 0, 2, 1, 2, 1}
// )

const (
	guessN = 22
	height = 16
)

var (
	gS = []string{"5616185650518293", "3847439647293047", "5855462940810587", "9742855507068353", "4296849643607543", "3174248439465858", "4513559094146117", "7890971548908067", "8157356344118483", "2615250744386899", "8690095851526254", "6375711915077050", "6913859173121360", "6442889055042768", "2321386104303845", "2326509471271448", "5251583379644322", "1748270476758276", "4895722652190306", "3041631117224635", "1841236454324589", "2659862637316867"}
	gI = []int{2, 1, 3, 3, 3, 1, 2, 3, 1, 2, 3, 1, 1, 2, 0, 2, 2, 3, 1, 3, 3, 2}
)

func known(info map[string]bool) (string, bool) {

	count := 0
	digit := ""

	for key, value := range info {
		if value {
			count++
			digit = key
		}

	}

	if count == 1 {
		return digit, true
	}

	return digit, false

}

func reSeed() (options []map[string]bool) {

	options = make([]map[string]bool, height)
	for i := 0; i < len(options); i++ {
		options[i] = make(map[string]bool)
		for j := 0; j < 10; j++ {
			digit := strconv.Itoa(j)
			options[i][digit] = true
		}
	}
	return
}

func length(state []map[string]bool) (lgth int) {
	for _, x := range state {

		for _, val := range x {
			if val {
				lgth++
			}
		}
	}
	return
}

//Does deduction on state in place
func infer(options []map[string]bool) {
	lgth := length(options)
	oldlgth := lgth + 10
	for oldlgth > lgth {

		for i := 0; i < guessN; i++ {
			correct := 0
			incorrect := 0

			for j := 0; j < height; j++ {

				place := gS[i][j : j+1]

				if digit, know := known(options[j]); know {
					if place == digit {
						correct++
					}
				}

				if !options[j][place] {
					incorrect++
				}
			}

			// fmt.Print(gS[i], " ", correct, incorrect)

			if correct == gI[i] {

				for j := 0; j < len(gS[i]); j++ {

					place := gS[i][j : j+1]

					if _, know := known(options[j]); !know {
						delete(options[j], place)
					}
				}
			}

			if incorrect+gI[i] == height {

				for j := 0; j < len(gS[i]); j++ {

					place := gS[i][j : j+1]

					if options[j][place] {

						options[j] = make(map[string]bool)
						options[j][place] = true

					}
				}
			}

			if correct > gI[i] || incorrect+gI[i] > height {
				options = make([]map[string]bool, 0)
				return
			}

			// fmt.Print("\n")
		}
		lgth, oldlgth = length(options), lgth
	}
}

func alive(options []map[string]bool) bool {
	for i := 0; i < len(options); i++ {
		if len(options[i]) == 0 {
			return false
		}
	}

	return true

}

func fixed(options []map[string]bool) bool {
	for i := 0; i < len(options); i++ {
		if len(options[i]) != 1 {
			return false
		}
	}

	return true
}

func guess(options []map[string]bool) {
	bGuessScore, bGuessIndex := 10000, 0

	for i := 0; i < len(options); i++ {
		if len(options[i]) > 1 && len(options[i]) < bGuessScore {
			bGuessIndex = i
			bGuessScore = len(options[i])
		}

	}

	pick := rand.Int() % bGuessScore
	guess := ""

	i := 0
	for key, _ := range options[bGuessIndex] {
		if i == pick {
			guess = key
		}
		i++
	}

	options[bGuessIndex] = make(map[string]bool)
	options[bGuessIndex][guess] = true
}

func stupid(options []map[string]bool) bool {

	for i := 0; i < guessN; i++ {
		correct := 0

		for j := 0; j < height; j++ {

			place := gS[i][j : j+1]

			if digit, know := known(options[j]); know {
				if place == digit {
					correct++
				}
			}

		}
		if correct != gI[i] {
			return true
		}
	}

	return false
}

func main() {
	starttime := time.Now()

	for i := 0; i < guessN; i++ {
		fmt.Println(gS[i], gI[i])
	}

	options := reSeed()

	for !fixed(options) {

		guess(options)

		infer(options)

		if (fixed(options) && stupid(options)) || !alive(options) {
			options = reSeed()
			infer(options)
		}

	}

	for i := 0; i < height; i++ {
		for dig, _ := range options[i] {
			fmt.Print(dig)
		}
	}
	fmt.Print("\n")

	fmt.Println("Elapsed time:", time.Since(starttime))
}