package main

import (
	"fmt"
	"time"
)

const N = 5
const circ = 1 << N

type hoop struct {
	val  [circ]bool
	seen map[[N]bool]bool
	fill int //index of first undetermined position
}

func construct() *hoop {
	answer := new(hoop)

	answer.seen = make(map[[N]bool]bool)

	tasty := [N]bool{}
	answer.seen[tasty] = true

	tasty[0] = true
	answer.seen[tasty] = true //backwards

	answer.fill = N
	return answer
}

func (a *hoop) possible() (works bool, which []bool) {
	test := [N]bool{}

	for i := 0; i < N-1; i++ {
		test[i] = a.val[a.fill-N+i+1]
	}

	which = make([]bool, 0)

	for _, test[N-1] = range []bool{true, false} {

		if !a.seen[test] {
			works = true
			which = append(which, test[N-1])
		}
	}

	return

}

func (a *hoop) valid() bool {

	for start := circ - N + 1; start < circ-1; start++ {
		test := [N]bool{}
		for i := 0; i < N; i++ {
			index := (start + i) % circ
			test[i] = a.val[index]
		}
		if a.seen[test] {
			return false
		}
		a.seen[test] = true
	}

	return true
}

func (a *hoop) append(toadd bool) (b *hoop) {
	b = new(hoop)
	b.val = a.val
	b.val[a.fill] = toadd
	b.fill = a.fill + 1
	b.seen = make(map[[N]bool]bool)
	for key, val := range a.seen {
		b.seen[key] = val
	}

	test := [N]bool{}
	for i := 0; i < N-1; i++ {
		test[i] = a.val[a.fill-N+i+1]
	}
	test[N-1] = toadd

	b.seen[test] = true
	return
}

func (a *hoop) evaluate() (answer int) {
	for i := 0; i < circ; i++ {
		if a.val[i] {
			answer += 1 << uint(circ-i-1)
		}
	}
	return
}

func main() {
	starttime := time.Now()

	var levels [circ - N + 1][]*hoop
	for i := 0; i < len(levels); i++ {
		levels[i] = make([]*hoop, 0)
	}

	levels[0] = append(levels[0], construct())

	for i := 0; i < len(levels)-1; i++ {

		for _, smaller := range levels[i] {

			works, options := smaller.possible()
			if works {
				for _, flag := range options {
					levels[i+1] = append(levels[i+1], smaller.append(flag))

				}

			}
		}

	}

	sum := int64(0)

	for _, hops := range levels[len(levels)-1] {

		if hops.valid() {
			sum += int64(hops.evaluate())
		}

	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
