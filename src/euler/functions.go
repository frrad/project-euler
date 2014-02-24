//Implement special mathematical functions

package euler

import (
//"fmt"
)

func TriangleNumber(n int) int64 {
	a := int64(n)
	return a * (a + 1) / 2
}
