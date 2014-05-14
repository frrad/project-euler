package main

import (
	"fmt"
	"strconv"
	"time"
)

func kill(input [][]int, I, J int) ([][]int, int) {
	ans := input[I][J]

	for i := I; i < len(input)-1; i++ {
		copy(input[i], input[i+1])
	}

	for i := 0; i < len(input); i++ {
		copy(input[i][J:len(input)-1], input[i][J+1:])
		input[i] = input[i][:len(input)-1]
	}

	return input[:len(input)-1], ans
}

func cans(input [][]int) [][2]int {
	ans := make([][2]int, 0)

	for i := 0; i < len(input); i++ {
		max, maxj := -1, -1
		for j := 0; j < len(input); j++ {
			if input[i][j] > max {
				max, maxj = input[i][j], j
			}
		}

		ans = append(ans, [2]int{i, maxj})
	}

	for i := 0; i < len(input); i++ {
		max, maxj := -1, -1
		for j := 0; j < len(input); j++ {
			if input[j][i] > max {
				max, maxj = input[j][i], j
			}
		}

		ans = append(ans, [2]int{maxj, i})
	}

	ans = dedupe(ans)

	return ans
}

func dedupe(input [][2]int) [][2]int {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				copy(input[i:len(input)-1], input[i+1:])
				input = input[:len(input)-1]
				i--
				break
			}
		}
	}

	return input
}

func trivial(couldBe [][2]int) [][2]int {
	ans := make([][2]int, 0)

	for _, can := range couldBe {
		i, j := can[0], can[1]
		x := 0
		for _, can2 := range couldBe {
			if can2[0] == i {
				x++
			}
			if can2[1] == j {
				x++
			}
		}
		if x == 2 {
			ans = append(ans, can)
		}
	}
	return ans
}

func trivialize(A [][]int) ([][]int, int) {
	total := 0

	for i := 0; i < len(A); i++ {

		couldBe := cans(A)
		crunch := trivial(couldBe)

		if len(crunch) != 0 {
			temp := 0
			i, j := crunch[0][0], crunch[0][1]
			A, temp = kill(A, i, j)
			total += temp
		}

	}

	return A, total
}

func clone(A [][]int) [][]int {
	size := len(A)
	B := make([][]int, size)
	for i := 0; i < size; i++ {
		B[i] = make([]int, size)
		copy(B[i], A[i])
	}
	return B
}

func keyFun(A [][]int) string {
	key := ""
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			key += strconv.Itoa(A[i][j])
			key += " "
		}
	}

	return key
}

var table = make(map[string]int)

func BFS(A [][]int) int {
	if len(A) == 0 {
		return 0
	}

	key := keyFun(A)

	if ans, ok := table[key]; ok {
		return ans
	}

	total := 0
	A, total = trivialize(A)

	best := -1
	for i := 0; i < len(A); i++ {
		B := clone(A)
		maxj, val := -1, -1
		for j := 0; j < len(A); j++ {
			if A[i][j] > val {
				val, maxj = A[i][j], j
			}
		}

		delta := 0
		B, delta = kill(B, i, maxj)
		downsize := delta + BFS(B)
		if downsize > best {
			best = downsize
		}

	}

	table[key] = total + best
	return total + best
}

func main() {
	starttime := time.Now()

	/*var A = [][]int{
		{7, 53, 183, 439, 863},
		{497, 383, 563, 79, 973},
		{287, 63, 343, 169, 583},
		{627, 343, 773, 959, 943},
		{767, 473, 103, 699, 303},
	}*/

	var A = [][]int{
		[]int{7, 53, 183, 439, 863, 497, 383, 563, 79, 973, 287, 63, 343, 169, 583},
		[]int{627, 343, 773, 959, 943, 767, 473, 103, 699, 303, 957, 703, 583, 639, 913},
		[]int{447, 283, 463, 29, 23, 487, 463, 993, 119, 883, 327, 493, 423, 159, 743},
		[]int{217, 623, 3, 399, 853, 407, 103, 983, 89, 463, 290, 516, 212, 462, 350},
		[]int{960, 376, 682, 962, 300, 780, 486, 502, 912, 800, 250, 346, 172, 812, 350},
		[]int{870, 456, 192, 162, 593, 473, 915, 45, 989, 873, 823, 965, 425, 329, 803},
		[]int{973, 965, 905, 919, 133, 673, 665, 235, 509, 613, 673, 815, 165, 992, 326},
		[]int{322, 148, 972, 962, 286, 255, 941, 541, 265, 323, 925, 281, 601, 95, 973},
		[]int{445, 721, 11, 525, 473, 65, 511, 164, 138, 672, 18, 428, 154, 448, 848},
		[]int{414, 456, 310, 312, 798, 104, 566, 520, 302, 248, 694, 976, 430, 392, 198},
		[]int{184, 829, 373, 181, 631, 101, 969, 613, 840, 740, 778, 458, 284, 760, 390},
		[]int{821, 461, 843, 513, 17, 901, 711, 993, 293, 157, 274, 94, 192, 156, 574},
		[]int{34, 124, 4, 878, 450, 476, 712, 914, 838, 669, 875, 299, 823, 329, 699},
		[]int{815, 559, 813, 459, 522, 788, 168, 586, 966, 232, 308, 833, 251, 631, 107},
		[]int{813, 883, 451, 509, 615, 77, 281, 613, 459, 205, 380, 274, 302, 35, 805},
	}

	fmt.Println(BFS(A))

	fmt.Println("Elapsed time:", time.Since(starttime))
}
