package euler

//Implement special mathematical functions

func TriangleNumber(n int) int64 {
	a := int64(n)
	return a * (a + 1) / 2
}
