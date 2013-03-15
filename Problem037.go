package main

import(
 "fmt"
 "math"
)

func isPrime(n int) bool{
lim := int(math.Sqrt(float64(n)))
for i:=2; i<lim ; i++{
if n% i ==0 {
return false
}
}

return true
}


func main() {

fmt.Println(isPrime(32259933))
}
