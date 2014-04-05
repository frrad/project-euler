package main

import (
	"fmt"
	"math/big"
	"time"
)

const players = 100

/*
        Grow    Stay   Shrink
Grow    1/36    1/9    1/36
Stay    1/9     4/9    1/9
Shrink  1/36    1/9    1/36

        Grow    Stay   Shrink
Grow    +2      +1     0
Stay    +1      0     -1
Shrink  0      -1     -2

+2 : 1/36
+1 : 2/9
 0 : 1/2
-1 : 2/9
-2 : 1/36
*/

func add(a, b int) int {
	return (a + b + players + players) % players
}

func transition(state *[players]*big.Rat) (won *big.Rat) {
	newState := new([players]*big.Rat)
	for i := 0; i < players; i++ {
		newState[i] = big.NewRat(0, 1)
	}

	tp := big.NewRat(1, 1)

	small, med, large := big.NewRat(1, 36), big.NewRat(2, 9), big.NewRat(1, 2)
	for i := 0; i < players; i++ {
		newState[add(i, -2)].Add(tp.Mul(small, state[i]), newState[add(i, -2)])
		newState[add(i, -1)].Add(tp.Mul(med, state[i]), newState[add(i, -1)])
		newState[add(i, 0)].Add(tp.Mul(large, state[i]), newState[add(i, 0)])
		newState[add(i, 1)].Add(tp.Mul(med, state[i]), newState[add(i, 1)])
		newState[add(i, 2)].Add(tp.Mul(small, state[i]), newState[add(i, 2)])
	}

	won = newState[0]
	newState[0] = big.NewRat(0, 1)

	for i := 0; i < players; i++ {
		state[i].Set(newState[i])
	}

	return
}

func main() {
	starttime := time.Now()

	state := new([players]*big.Rat)
	for i := 0; i < players; i++ {
		state[i] = big.NewRat(0, 1)
	}
	state[players/2].SetFloat64(1)

	dead := big.NewRat(0, 1)
	expected := big.NewRat(0, 1)

	for i := 1; ; i++ {
		here := transition(state)

		dead.Add(dead, here)
		move := big.NewRat(int64(i), 1)
		expected.Add(expected, move.Mul(here, move))

		fmt.Printf("(%d) Total Dead: %s\t This Turn: %s\t Expected Length: %s\n", i, dead.FloatString(4), here.FloatString(4), expected.FloatString(10))

		//fmt.Printf("%s\n", state)

	}

	fmt.Println("Elapsed time:", time.Since(starttime))
}
