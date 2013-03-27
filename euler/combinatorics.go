package euler

//returns the nth permutation of the given slice
func Permutation(n int, list []int) []int {
	if len(list) == 1 {
		return list
	}

	k := n % len(list)

	first := []int{list[k]}
	next := make([]int, len(list)-1)

	copy(next, append(list[:k], list[k+1:]...))

	return append(first, Permutation(n/len(list), next)...)

}

func PermuteString(n int, word string) string {
	if len(word) == 1 {
		return word
	}

	k := n % len(word)

	return word[k:k+1] + PermuteString(n/len(word), word[:k]+word[k+1:])
}

//returns which permutation takes a->b, or -1
//NOTE: THIS IS A TERRIBLE ALGORITHM -- Fix later
func UnPermuteStrings(a, b string) int {
	for i := 0; i < int(Factorial(int64(len(a)))); i++ {
		if PermuteString(i, a) == b {
			return i
		}
	}
	return -1

}
