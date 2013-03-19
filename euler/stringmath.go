package eulerlib

import (
	"strconv"
)

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
