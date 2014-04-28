package main

import (
	"fmt"
	"time"
)

const tip = 8

func above(line []int) int64 {
	if len(line) == 3 {
		return 2
	}

	if ans, ok := memo[keyFun(line)]; ok {
		return ans
	}

	ans := int64(0)
	cans := candidates(line)

	for _, can := range cans {
		ans += above(can)
	}

	memo[keyFun(line)] = ans
	return ans
}

func candidates(template []int) [][]int {
	ans := make([][]int, 0)
	avoid := avoid(template)
	bits := uint((len(avoid) + 1) / 2)

	for i := 0; i < 1<<bits; i++ {
		mask := make([]int, len(avoid))
		copy(mask, avoid)
		mspec := i
		for j := uint(0); j < bits; j++ {
			tap := mspec % 2
			mspec /= 2
			mask[2*j] += 1 + tap
			mask[2*j] -= 1
			mask[2*j] %= 3
			mask[2*j] += 1
		}

		tocolor := make([][]int, 0)

		for fill := 1; fill < len(mask); fill += 2 {
			if mask[fill-1] == mask[fill+1] {
				tocolor = append(tocolor, []int{fill, mask[fill-1]})
			} else {
				mask[fill] = 6 - mask[fill-1] - mask[fill+1]
			}
		}

		slots := uint(len(tocolor))
		for j := 0; j < 1<<slots; j++ {
			info := j
			clink := make([]int, len(mask))
			copy(clink, mask)
			for k := uint(0); k < slots; k++ {
				pos := info % 2
				info /= 2
				toput := tocolor[k][1]
				toput = (toput+pos)%3 + 1
				clink[tocolor[k][0]] = toput
			}
			ans = append(ans, clink)
		}

	}
	return ans
}

func avoid(line []int) []int {
	ans := make([]int, len(line)-2)
	for i := range ans {
		if i%2 == 0 {
			ans[i] = line[i+1]
		}
	}
	return ans
}

var memo = make(map[int64]int64)

func keyFun(line []int) (key int64) {
	for _, val := range line {

		key *= 3 // base 4 rep
		key += int64(val) - 1
	}
	return
}

func main() {
	starttime := time.Now()

	accumulate := int64(0)

	top := uint(tip*2 - 1)

	test := make([]int, top)
	test[0] = 1

	for i := uint(0); i < 1<<(top-1); i++ {
		key := i
		for j := uint(1); j < top; j++ {
			mod := key % 2
			key /= 2
			test[j] = int((uint(test[j-1])+mod)%3 + 1)

		}
		//		fmt.Println(test)
		accumulate += above(test)
	}

	fmt.Println(accumulate * 3)

	fmt.Println("Elapsed time:", time.Since(starttime))

}
