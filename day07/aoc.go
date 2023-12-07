package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Hand struct {
	cards string
	bid   uint64
	order uint64
}

func getSolutionPart1(input string) uint64 {
	hands := []Hand{}
	strength := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}
		h := strings.Split(row, " ")
		cards := h[0]
		bid, _ := strconv.ParseUint(h[1], 10, 64)

		counts := make(map[rune]int)
		for _, char := range cards {
			counts[char]++
		}

		hand := Hand{cards, bid, 0}

		switch len(counts) {
		case 1:
			hand.order = 1

		case 2:
			for _, v := range counts {
				if v == 4 {
					hand.order = 2
					break
				}
				if v == 3 {
					hand.order = 3
					break
				}
			}

		case 3:
			for _, v := range counts {
				if v == 3 {
					hand.order = 4
					break
				}
				if v == 2 {
					hand.order = 5
					break
				}
			}
		case 4:
			hand.order = 6

		case 5:
			hand.order = 7
		}

		hands = append(hands, hand)

	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].order == hands[j].order {
			x := 0
			for hands[i].cards[x] == hands[j].cards[x] {
				x++
			}
			return strength[rune(hands[i].cards[x])] < strength[rune(hands[j].cards[x])]
		}
		return hands[i].order > hands[j].order
	})

	result := uint64(0)
	i := uint64(1)
	for _, hand := range hands {
		result += hand.bid * i
		i++
	}

	return result
}

func getSolutionPart2(input string) uint64 {
	hands := []Hand{}
	strength := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 11,
		'9': 10,
		'8': 9,
		'7': 8,
		'6': 7,
		'5': 6,
		'4': 5,
		'3': 4,
		'2': 3,
		'J': 2,
	}

	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}
		h := strings.Split(row, " ")
		cards := h[0]
		bid, _ := strconv.ParseUint(h[1], 10, 64)

		counts := make(map[rune]int)
		jokerCount := 0
		for _, char := range cards {
			if char == 'J' {
				jokerCount++
			} else {
				counts[char]++
			}
		}

		hand := Hand{cards, bid, 0}

		switch len(counts) {
		case 0:
			hand.order = 1
		case 1:
			hand.order = 1

		case 2:
			for _, v := range counts {
				if v+jokerCount == 4 {
					hand.order = 2
					break
				}
				if v+jokerCount == 3 {
					hand.order = 3
				}
			}

		case 3:
			for _, v := range counts {
				if v+jokerCount == 3 {
					hand.order = 4
					break
				}
				if v+jokerCount == 2 {
					hand.order = 5
				}
			}
		case 4:
			hand.order = 6
		case 5:
			hand.order = 7
		}

		hands = append(hands, hand)

	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].order == hands[j].order {
			x := 0
			for hands[i].cards[x] == hands[j].cards[x] {
				x++
			}
			return strength[rune(hands[i].cards[x])] < strength[rune(hands[j].cards[x])]
		}
		return hands[i].order > hands[j].order
	})

	result := uint64(0)
	i := uint64(1)
	for _, hand := range hands {
		result += hand.bid * i
		i++
	}

	return result
}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
