package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	breedSize = 100
	genSize   = breedSize * breedSize
	keepers   = 95
	infinity  = 3000
	churn     = 10 //How long do we refine in PP?
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

func random() (r string) {
	for i := 0; i < height; i++ {
		r += strconv.Itoa(rand.Int() % 10)
	}
	return
}

func score(a string) (an int) {
	for i := 0; i < guessN; i++ {
		count := 0
		for j := 0; j < height; j++ {
			if a[j] == gS[i][j] {
				count++
			}
		}
		// fmt.Println(gS[i], a, count)
		an += abs(count - gI[i])
	}
	return
}

func insert(a string, set []string) int {
	index := 0

	for ; index < len(set) && score(set[index]) < score(a); index++ {
	}

	if index < len(set)-2 {
		copy(set[index+2:], set[index+1:len(set)-2])
	}

	set[index] = a

	return score(set[len(set)-1])

}

func abs(an int) int {
	if an < 0 {
		return -1 * an
	}
	return an
}

func sort(list []string) int {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if score(list[j]) > score(list[j+1]) {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}

	return score(list[len(list)-1])

}

func mix(a, b string) (mash string) {
	for i := 0; i < height; i++ {
		if rand.Int()%2 == 0 {
			mash += a[i : i+1]
		} else {
			mash += b[i : i+1]
		}
	}
	// fmt.Println(a, b, "=>", mash)
	return
}

func evolve(termAge int, generation [genSize]string) string {
	answer := generation[0]
	brake := false

	for age := 0; !brake && age != termAge; age++ {
		//Prep winners set with first entries
		winners := make([]string, keepers)
		for i := 0; i < keepers; i++ {
			winners[i] = generation[i]
		}
		threshold := sort(winners)

		//Add best to winners
		for i := keepers; i < genSize; i++ {
			if sc := score(generation[i]); sc < threshold {
				answer = generation[i]
				if sc == 0 {
					brake = true
					break
				}
				threshold = insert(generation[i], winners)
			}
		}

		fmt.Print("\r", winners[0], "      ", score(winners[0]), "      ", score(winners[keepers-1]), "       ")

		breed := [breedSize]string{}

		for i := 0; i < keepers; i++ {
			breed[i] = winners[i]
		}

		for i := keepers; i < breedSize; i++ {
			breed[i] = random()
		}

		for i := 0; i < breedSize; i++ {
			for j := i + 1; j < breedSize; j++ {
				if breed[i] == breed[j] {
					breed[i] = random()
				}
			}
		}

		for i := 0; i < breedSize; i++ {
			for j := 0; j < breedSize; j++ {
				generation[i*breedSize+j] = mix(breed[i], breed[j])
			}
		}

	}

	return answer

}

func main() {
	starttime := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())

	var load [genSize]string

	if churn != 0 {
		var smart [breedSize]string

		for i := 0; i < breedSize; i++ {
			var generation [genSize]string
			//Seed Population
			for i := 0; i < genSize; i++ {
				generation[i] = random()
			}

			smart[i] = evolve(churn, generation)

			fmt.Println("\tLoaded", i+1)

			if score(smart[i]) == 0 {
				fmt.Println("Answer found early:\n", smart[i])
				panic("")
			}
		}

		for i := 0; i < breedSize; i++ {
			for j := 0; j < breedSize; j++ {
				load[i*breedSize+j] = mix(smart[i], smart[j])
			}
		}

	} else {
		for i := 0; i < genSize; i++ {
			load[i] = random()
		}
	}

	winning := evolve(-1, load)
	fmt.Println("\n", winning)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
