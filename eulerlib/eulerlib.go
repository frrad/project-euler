package eulerlib

import (
	"math"
	"strconv"
)

const (
	primeTableLength = 100000000
	//lastPrime = Prime[primeTableLength - 1]
	lastPrime = 2038074739
)

var primeTable [primeTableLength]int64

func Prime(n int64) int64 {

	if n < 1 {
		return 0
	}

	primeTable[1] = 2
	primeTable[2] = 3

	if n < primeTableLength && primeTable[n] != 0 {
		return primeTable[n]
	}

	i := Prime(n-1) + 1

	for !IsPrime(i) {
		i++
	}

	if n < primeTableLength {
		primeTable[n] = i
	}
	return i
}

func IsPrime(n int64) bool {

	end := int64(math.Sqrt(float64(n)))

	//If we start computing beyond the table, this is stupid
	for i := int64(1); Prime(i) <= end && i < primeTableLength; i++ {
		if n%Prime(i) == 0 {
			return false
		}
	}

	//If we need to pass the end of the prime table brute force
	if end > lastPrime {
		for i := int64(lastPrime); i <= end; i++ {
			if n%i == 0 {
				return false
			}
		}

	}

	return true
}

func ArePermutations(a int64, b int64) bool {
	A := strconv.FormatInt(a, 10)
	B := strconv.FormatInt(b, 10)

	length := len(A)
	list1 := make([]byte, length)
	list2 := make([]byte, length)

	if len(A) != len(B) {
		return false
	}

	for i := 0; i < length; i++ {
		list1[i] = A[i]
		list2[i] = B[i]
	}

	for i := 0; i < length; i++ {
		flag := false

		for j := 0; j < length; j++ {
			if flag == false && list1[i] == list2[j] {
				list2[j] = 0
				flag = true
			}

		}
		if flag == false {
			return false
		}
	}
	return true

}

func Min(m int64, n int64) int64 {
	if m < n {
		return m
	}
	return n
}

func GCD(n int64, m int64) int64 {

	top := Min(m, n)

	for i := int64(1); Prime(i) <= top; i++ {
		if n%Prime(i) == 0 && m%Prime(i) == 0 {
			return Prime(i) * GCD(n/Prime(i), m/Prime(i))
		}
	}
	return 1
}

func FracAdd(num1 int64, den1 int64, num2 int64, den2 int64) (num int64, den int64) {
	gcd := GCD(den1, den2)
	d1 := den1 / gcd
	d2 := den2 / gcd

	den = d1 * d2 * gcd
	num = num1*d2 + num2*d1

	return
}

func FracReduce(num int64, den int64) (int64, int64) {
	gcd := GCD(num, den)
	return num / gcd, den / gcd
}

func DigitSum(N int64) (sum int) {

	n := int(N)
	sum = 0
	word := strconv.Itoa(n)

	for i := 0; i < len(word); i++ {
		x, _ := strconv.Atoi(string(word[i]))
		sum += x
	}
	return
}

func Max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func StringReverse(a string) string {
	b := ""
	for i := len(a) - 1; i >= 0; i-- {
		b += string(a[i])
	}
	return b
}

func IntReverse(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	s = StringReverse(s)
	x, _ := strconv.ParseInt(s, 10, 64)
	return x
}

func IsPalindrome(n int64) bool {
	if n == IntReverse(n) {
		return true
	}
	return false
}

func IsStringPalindrome(n string) bool {
	if n == StringReverse(n) {
		return true
	}
	return false
}

//Removes 0-padding on Integer Strings
func StringTrim(a string) string {
	if a == "0" {
		return a
	}
	place := 0
	for i := 0; i < len(a) && string(a[i]) == "0"; i++ {
		place = i + 1
	}

	output := ""

	for i := place; i < len(a); i++ {
		output += string(a[i])
	}

	return output
}

func StringSum(string1 string, string2 string) string {
	length1 := int64(len(string1))
	length2 := int64(len(string2))
	length := 1 + Max(length1, length2)
	string1 = StringReverse(string1)
	string2 = StringReverse(string2)

	sum := make([]int, length)
	str1 := make([]int, length)
	str2 := make([]int, length)

	for i := int64(0); i < length; i++ {
		a := 0
		b := 0
		sum[i] = 0

		if i < length1 {
			a, _ = strconv.Atoi(string(string1[i]))

		}
		str1[i] = a

		if i < length2 {
			b, _ = strconv.Atoi(string(string2[i]))

		}

		str2[i] = b

	}

	for i := int64(0); i < length-1; i++ {
		total := str1[i] + str2[i] + sum[i]
		sum[i] = total % 10
		sum[i+1] = (total - total%10) / 10
	}

	answer := ""

	for i := int64(0); i < length; i++ {
		answer += strconv.Itoa(sum[i])
	}

	answer = StringReverse(answer)
	answer = StringTrim(answer)

	return answer
}
