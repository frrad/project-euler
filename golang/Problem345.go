package main

import (
	"fmt"
	"time"
)

//Use Hungarian algorithm to minimize  - A

func print(an [][]int) {
	fmt.Println("")
	for i := 0; i < len(an); i++ {
		for _, p := range an[i] {
			fmt.Print(p, "\t")
		}
		fmt.Print("\n")
	}
}

func reverse(a [][]int) {
	for _, row := range a {
		for i := range row {
			row[i] = -row[i]
		}
	}
}

func max(A [][]int) int {
	mix := A[0][0]
	for _, row := range A {
		for i := range row {
			if row[i] > mix {
				mix = row[i]
			}
		}
	}
	return mix
}

func matAdd(A [][]int, alpha int) {
	for _, row := range A {
		for i := range row {
			row[i] = alpha + row[i]
		}
	}
}

func rowMin(A [][]int, i int) int {
	row := A[i]
	min := row[0]
	for _, fig := range row {
		if fig < min {
			min = fig
		}
	}
	return min
}

func rowAdd(A [][]int, i int, alpha int) {
	row := A[i]
	for j := range row {
		row[j] += alpha
	}
}

func transpose(A [][]int) {
	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			A[i][j], A[j][i] = A[j][i], A[i][j]
		}
	}
}

func minor(A [][]int, a, b int) {
	A = append(A[:a], A[a+1:]...)
	for _, row := range A {

		row = append(row[:b], row[b+1:]...)
	}

}

func reduce(A [][]int) [][]int {

	lid := len(A)

	for x := 0; x < len(A); x++ {

		for i := 0; i < lid; i++ {
			for j := 0; j < lid && i < lid; j++ {
				//	fmt.Println(i, j, lid)
				//	print(A)
				if A[i][j] == 0 {
					minor(A, i, j)
					lid--
				}
			}

		}

	}

	B := make([][]int, lid)
	for i := 0; i < lid; i++ {
		B[i] = make([]int, lid)
		copy(B[i], A[i][:lid])
	}
	return B

}
func main() {
	starttime := time.Now()

	var A = [][]int{
		[]int{7, 53, 183, 439, 863},
		[]int{497, 383, 563, 79, 973},
		[]int{287, 63, 343, 169, 583},
		[]int{627, 343, 773, 959, 943},
		[]int{767, 473, 103, 699, 303},
	}

	/*	var Aaaaa = [][]int{
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
	}*/

	max := max(A)
	reverse(A)
	matAdd(A, max)
	print(A)

	pay := 0
	for i := 0; i < len(A); i++ {
		twit := rowMin(A, i)
		pay += twit
		rowAdd(A, i, -twit)
	}

	print(A)

	transpose(A)
	for i := 0; i < len(A); i++ {
		twit := rowMin(A, i)
		pay += twit
		rowAdd(A, i, -twit)
	}

	transpose(A)
	print(A)

	fmt.Println(len(A)*max - pay)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
