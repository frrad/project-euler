package main

import (
	"./euler"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type card struct {
	value int
	suit  string
}

func makeCard(letters string) card {
	value, err := strconv.Atoi(letters[:1])
	if err != nil {
		if letters[:1] == "T" {
			value = 10
		}
		if letters[:1] == "J" {
			value = 11
		}
		if letters[:1] == "Q" {
			value = 12
		}
		if letters[:1] == "K" {
			value = 13
		}
		if letters[:1] == "A" {
			value = 14
		}
	}

	return card{value, letters[1:]}
}

func hasPair(hand []card) bool {
	for i, card := range hand {
		for _, match := range hand[i+1:] {
			if match.value == card.value {
				return true
			}
		}
	}
	return false
}

func has3(hand []card) bool {
	for i, card1 := range hand {
		for j, card2 := range hand[i+1:] {
			for _, card3 := range hand[i+j+2:] {
				if card1.value == card2.value && card2.value == card3.value {
					fmt.Println(card1.value)
					return true

				}
			}
		}
	}
	return false
}

func hasflush(hand []card) bool {
	suit := hand[0].suit
	for _, card := range hand {
		if suit != card.suit {
			return false
		}
	}
	return true

}

func main() {
	starttime := time.Now()

	data := euler.Import("problemdata/poker.txt")

	for _, line := range data {

		hand1 := make([]card, 5)
		for i, card := range strings.Split(line[:14], " ") {
			hand1[i] = makeCard(card)
		}

		hand2 := make([]card, 5)
		for i, card := range strings.Split(line[15:], " ") {
			hand2[i] = makeCard(card)
		}

		if hasflush(hand1) {
			fmt.Println(hand1, hand2)

		}

	}

	fmt.Println("Elapsed time:", time.Since(starttime))

}
