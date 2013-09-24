package euler

//returns the square of a matrix
func SqrIntMatrix(A [][]int) [][]int {
	n := len(A)
	square := make([][]int, n)
	for i := 0; i < n; i++ {
		square[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			answer := 0
			for k := 0; k < n; k++ {
				answer += A[j][k] * A[k][i]
			}
			square[j][i] = answer
		}
	}

	return square
}

func MatrixProd(a, b [][]float64) [][]float64 {
	i := len(a)
	j := len(a[0])
	if len(b) != j {
		panic("matrices mismatched")
	}
	k := len(b[0])

	prod := make([][]float64, i)

	for row := 0; row < i; row++ {
		theRow := make([]float64, k)
		for column := 0; column < k; column++ {
			elt := 0.
			for index := 0; index < j; index++ {
				elt += a[row][index] * b[index][column]
			}
			theRow[column] = elt
		}
		prod[row] = theRow
	}

	return prod
}
