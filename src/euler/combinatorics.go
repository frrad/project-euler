package euler

import "math/big"

var parttable = make(map[int]int)

//IntPartitions returns the number of partitions of n objects into groups of 1
//or more. Currently implemented using a recurrence relation, due to Euler.
func IntPartitions(n int) int {
	if ans, ok := parttable[n]; ok {
		return ans
	}

	if n == 0 {
		parttable[0] = 1
		return IntPartitions(0)
	}

	if n < 0 {
		return 0
	}

	sum := 0

	for k := 1; k <= n; k++ {
		var summand int
		if k%2 == 0 {
			summand = -1
		} else {
			summand = 1
		}

		summand *= IntPartitions(n-(k*(3*k-1)/2)) + IntPartitions(n-(k*(3*k+1)/2))

		sum += summand
	}

	parttable[n] = sum
	return IntPartitions(n)
}

var sfcache = make(map[int]int64)

//SubFactorial(n) gives the subfactorial of n (i.e. !n). This is also the number
//of permutations of n elements, fixing none of them (derangements). The output
//starts to overflow near !20 or !21. (See BigSubFactorial for big.Int
//implementation.
func Subfactorial(n int) int64 {
	if n == 0 {
		return 1
	}

	if n <= 1 {
		return 0
	}

	if ans, ok := sfcache[n]; ok {
		return ans
	}

	ans := int64(n-1) * (Subfactorial(n-1) + Subfactorial(n-2))
	sfcache[n] = ans

	return Subfactorial(n)
}

var bigsfcache = make(map[int]*big.Int)

//BigSubfactorial is the same as SubFactorial, but returns a big.Int object to
//avoid overflow issues.
func BigSubfactorial(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	if n <= 1 {
		return big.NewInt(0)
	}

	if ans, ok := bigsfcache[n]; ok {
		//make a copy so user can't modify cache
		ret := new(big.Int)
		return ret.Set(ans)
	}

	d1 := BigSubfactorial(n - 1)
	d2 := BigSubfactorial(n - 2)
	n1 := big.NewInt(int64(n - 1))

	d1.Add(d1, d2)
	bigsfcache[n] = n1.Mul(d1, n1)

	return BigSubfactorial(n)
}

var bigChooseCache = make(map[[2]int]*big.Int)

//BigChoose(n,k) is equivalent to "n choose k." Binomial coefficients are
//currently computed using a recursive, "Pascal's Triangle" approach (with
//memoization).
func BigChoose(n, k int) *big.Int {
	if k == 0 || k == n {
		return big.NewInt(1)
	}
	if k < 0 || k > n {
		return big.NewInt(0)
	}

	if ans, ok := bigChooseCache[[2]int{n, k}]; ok {
		//make a copy so user can't modify cache
		ret := new(big.Int)
		return ret.Set(ans)
	}

	ans := BigChoose(n-1, k)
	ans.Add(ans, BigChoose(n-1, k-1))

	bigChooseCache[[2]int{n, k}] = ans
	return BigChoose(n, k)
}

//Choose(n,k) gives "n choose k." Currently implemented using a version of the
//factorial formula, but working with prime decompositions to avoid integer
//overflow in intermediate steps.
func Choose(N, K int64) int64 {
	factors := make(map[int64]int64)

	if K == 0 || N == K || N <= 1 {
		return 1
	}

	if N < K {
		return 0
	}

	for n := N; n > N-K; n-- {
		nfactors := Factors(n)
		for i := 0; i < len(nfactors); i++ {
			factors[nfactors[i][0]] += nfactors[i][1]
		}
	}

	for k := K; k >= 2; k-- {

		kfactors := Factors(k)
		for i := 0; i < len(kfactors); i++ {
			factors[kfactors[i][0]] -= kfactors[i][1]
		}

	}

	answer := int64(1)

	for prime, multiplicity := range factors {

		for i := int64(0); i < multiplicity; i++ {
			answer *= prime
		}

	}
	return answer
}

//Permutation returns the nth permutation of a slice of integer
//values. Undefined behavior for n > (len(list) factorial).
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

//PermuteFloats is just like Permutation, but for slices of float64 values. It
//returns the nth permutation of the given slice.
func PermuteFloats(n int, list []float64) []float64 {
	if len(list) == 1 {
		return list
	}

	k := n % len(list)

	first := []float64{list[k]}
	next := make([]float64, len(list)-1)

	copy(next, append(list[:k], list[k+1:]...))

	return append(first, PermuteFloats(n/len(list), next)...)
}

//PermuteString returns the nth permutation of the supplied string
//(cf. Permutation, and PermuteFloats).
func PermuteString(n int, word string) string {
	if len(word) == 1 {
		return word
	}

	k := n % len(word)

	return word[k:k+1] + PermuteString(n/len(word), word[:k]+word[k+1:])
}

func SplitInts(list []int, K, N int) (a, b []int) {
	a, b = make([]int, 0), make([]int, 0)

	indices := make(map[int]bool)

	for k := K; k >= 1; k-- {

		n := k - 1

		if Choose(int64(n), int64(k)) <= int64(N) {
			for ; Choose(int64(n), int64(k)) <= int64(N); n++ {

			}
			n--
		}

		indices[n] = true

		N = N - int(Choose(int64(n), int64(k)))
	}

	a, b = make([]int, 0), make([]int, 0)

	for i := 0; i < len(list); i++ {
		if indices[i] {
			a = append(a, list[i])
		} else {
			b = append(b, list[i])
		}
	}

	return a, b
}

func SplitSeq(K, N int) (a []int) {

	indices := make([]int, 0)

	for k := K; k >= 1; k-- {

		n := k - 1

		if Choose(int64(n), int64(k)) <= int64(N) {
			for ; Choose(int64(n), int64(k)) <= int64(N); n++ {

			}
			n--
		}

		indices = append(indices, n)

		N = N - int(Choose(int64(n), int64(k)))
	}

	return indices
}

//UnPermuteStrings returns the index of the first permutation which takes a to
//b. That is, it finds n which has PermuteString(n, a) == b. Returns -1 on
//failure. Current implementation is just brute force: could be substantially
//improved if necessary.
func UnPermuteStrings(a, b string) int {
	for i := 0; i < int(Factorial(int64(len(a)))); i++ {
		if PermuteString(i, a) == b {
			return i
		}
	}
	return -1

}
