package main

import (
	"euler"
	"fmt"
	"strconv"
	"strings"
)

const (
	path   = "../eulerdata/status.html"
	prizes = 5
)

var totals [prizes]int

var names = [prizes]string{
	"Prime Obsession",
	"Triangle Trophy",
	"Lucky Luke\t",
	"Decimation II\t",
	"Ultimate Decimator",
}

var taglines = [prizes]string{
	"prime numbered problems",
	"first triangle numbered problems",
	"lucky numbered problems",
	"rows",
	"rows",
}

var goals = [prizes]int{
	50,
	25,
	50,
	10,
	10,
}

var prizeFns = make([]func(map[int]bool) (int, map[int]bool), prizes)

func show(set map[int]bool, howHard map[int]int) string {
	ans := ""
	for i := 0; i < max; i++ {
		if set[i] {
			ans += strconv.Itoa(i)
			ans += "("
			ans += strconv.Itoa(howHard[i])
			ans += ") "
		}
	}
	return ans
}

func getNum(a string) int {
	probLen := 8 //Length of `Problem '

	starts := strings.Index(a, "Problem ")
	ends := strings.Index(a[starts+probLen:], " ")

	isolated := a[starts+probLen : starts+probLen+ends]
	number, _ := strconv.Atoi(isolated)

	return number
}

func luckySeive(max int) []int {
	luckyseive := make([]int, max)
	for i := 0; i < max; i++ {
		luckyseive[i] = i + 1
	}

	last := -1
	pointer := 1

	for pointer < len(luckyseive) {

		last = luckyseive[pointer]

		for del := last - 1; del < len(luckyseive); del += last {
			luckyseive[del] = 0
		}
		for i := 0; i < len(luckyseive); i++ {
			if luckyseive[i] == 0 {
				luckyseive = append(luckyseive[:i], luckyseive[i+1:]...)
				i--
			}

		}

		if luckyseive[pointer] == last {
			pointer++
		}

	}

	return luckyseive
}

func howHard(text string) int {
	start := strings.Index(text, "solved by")
	start += 10 //length of "solved by"
	text = text[start:]

	end := strings.Index(text, "members")
	text = text[:end-1]

	//fmt.Printf("%s\n\n\n", text)

	ans, err := strconv.Atoi(text)
	if err == nil {
		//fmt.Printf("$d\n", ans)
		return ans
	}

	fmt.Printf("ERROR: %s\n", err)
	return 0
}

var max int = -1 //number of problems total

func main() {

	prizeFns[0] = func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		//PRIME NUMBERS (Index = 0)
		for i := 1; i <= max; i++ {
			if dict[i] {
				if euler.IsPrime(int64(i)) {
					ans++
				}
			} else if euler.IsPrime(int64(i)) {
				set[i] = true
			}
		}
		return
	}

	prizeFns[1] = func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		//TRIANGLE NUMBERS (Index = 1)
		for i := 1; i <= 25; i++ {
			if dict[i*(i+1)/2] {
				ans++
			} else {
				set[i*(i+1)/2] = true
			}
		}
		return
	}

	prizeFns[2] = func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		//LUCKY NUMBER (Index = 2)
		luckyseive := luckySeive(max)
		for i := 0; i < len(luckyseive); i++ {
			if dict[luckyseive[i]] {
				ans++
			} else {
				set[luckyseive[i]] = true
			}
		}
		return
	}

	prizeFns[3] = func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		//DECIMATION II (Index = 3)
		decStart := 200
		for i := 0; i < 10; i++ {
			here := 0
			for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
				if dict[j] {
					here++
				}
			}

			if here > 0 {
				ans++
			} else {

				for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
					set[j] = true
				}
			}
		}
		return
	}

	prizeFns[4] = func(dict map[int]bool) (ans int, set map[int]bool) {
		set = make(map[int]bool)
		//ULTIMATE DECIMATOR (Index = 4)
		decStart := 300
		for i := 0; i < 10; i++ {
			here := 0
			for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
				if dict[j] {
					here++
				}
			}

			if here > 0 {
				ans++
			} else {

				for j := decStart + 10*i + 1; j < decStart+10*(i+1)+1; j++ {
					set[j] = true
				}
			}
		}
		return
	}

	lineL := 40

	page := euler.Import(path)

	dict := make(map[int]bool)
	difficulty := make(map[int]int)

	for _, line := range page {
		split := strings.Split(line, "vertical-align:middle;background-color:#")
		for _, prob := range split {

			if len(prob) > 5 {
				if prob[:6] == "CEE7B6" {
					//Green = Complete
					number := getNum(prob)
					difficulty[number] = howHard(prob)
					dict[number] = true
					if number > max {
						max = number
					}

				} else if prob[:3] == "fff" {
					//White = Incomplete
					number := getNum(prob)
					difficulty[number] = howHard(prob)
					if number > max {
						max = number
					}
				}
			}

		}
	}

	done := 0
	for i := 1; i <= max; i++ {
		if dict[i] {
			done++
		}
	}

	for i := 0; i < prizes; i++ {
		totals[i], _ = prizeFns[i](dict)
	}

	fmt.Println("Done", done, "/", max, " problems")

	for i := 0; i < lineL; i++ {
		fmt.Print("=")
	}
	fmt.Print("\n")

	for i := 1; i <= max; i++ {
		if dict[i] {
			fmt.Print("X")
		} else {
			fmt.Print(" ")
		}

		if i%lineL == 0 {
			fmt.Print("\n")
		}
	}
	if max%lineL != 0 {
		fmt.Print("\n")
	}
	for i := 0; i < lineL; i++ {
		fmt.Print("=")
	}
	fmt.Print("\n\n")

	for i := 0; i < prizes; i++ {
		fmt.Printf("%s \t %d/%d %s\n", names[i], totals[i], goals[i], taglines[i])
	}

	fmt.Print("\n")

	track := make(map[int]int)

	for pNum := 1; pNum <= 4; pNum++ {

		if totals[pNum] < goals[pNum] {
			_, set := prizeFns[pNum](dict)
			fmt.Printf("%s: %s\n", names[pNum], show(set, difficulty))

			for i, _ := range set {
				track[i]++
			}

		}
	}

	maxTrack := -1

	fmt.Print("\nRepeats: ")

	for i := 1; i <= max; i++ {
		if track[i] > 1 {
			fmt.Print(i, " ")

		}
		if track[i] > maxTrack {
			maxTrack = track[i]
		}
	}

	fmt.Print("\n")

	set := make(map[int]bool)
	for i := 1; i <= max; i++ {
		if track[i] == maxTrack {
			set[i] = true
		}
	}
	fmt.Printf("Most Repeated (%d): %s", maxTrack, show(set, difficulty))

	fmt.Print("\n")

}
