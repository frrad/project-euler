package main

import (
	"fmt"
	"time"
)

func spell(n int) string {
	if n == 1 {
		return "one"
	}
	if n == 2 {
		return "two"
	}
	if n == 3 {
		return "three"
	}
	if n == 4 {
		return "four"
	}
	if n == 5 {
		return "five"
	}
	if n == 6 {
		return "six"
	}
	if n == 7 {
		return "seven"
	}
	if n == 8 {
		return "eight"
	}
	if n == 9 {
		return "nine"
	}
	if n == 10 {
		return "ten"
	}
	if n == 11 {
		return "eleven"
	}
	if n == 12 {
		return "twelve"
	}
	if n == 13 {
		return "thirteen"
	}
	if n == 14 {
		return "fourteen"
	}
	if n == 15 {
		return "fifteen"
	}
	if n == 16 {
		return "sixteen"
	}
	if n == 17 {
		return "seventeen"
	}
	if n == 18 {
		return "eighteen"
	}
	if n == 19 {
		return "nineteen"
	}

	if n > 19 && n < 100 {
		tens := n / 10
		ones := n % 10
		if tens == 2 {
			return "twenty " + spell(ones)
		}
		if tens == 3 {
			return "thirty " + spell(ones)
		}
		if tens == 4 {
			return "forty " + spell(ones) //not spelled fourty!
		}
		if tens == 5 {
			return "fifty " + spell(ones)
		}
		if tens == 6 {
			return "sixty " + spell(ones)
		}
		if tens == 7 {
			return "seventy " + spell(ones)
		}
		if tens == 8 {
			return "eighty " + spell(ones)
		}
		if tens == 9 {
			return "ninety " + spell(ones)
		}
	}

	if n > 99 && n < 1000 {
		ending := n % 100
		hundreds := n / 100

		var endSpell string
		if ending == 0 {
			endSpell = ""
		} else {
			endSpell = " and " + spell(ending)
		}

		return spell(hundreds) + " hundred" + endSpell

	}

	if n == 1000 {
		return "one thousand"
	}

	return ""

}

func count(word string) int {
	if word == "" {
		return 0
	}
	if string(word[0]) == " " {
		return count(word[1:])
	}
	return 1 + count(word[1:])
}

func main() {
	starttime := time.Now()

	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += count(spell(i))
	}

	fmt.Println(sum)

	fmt.Println("Elapsed time:", time.Since(starttime))
}
