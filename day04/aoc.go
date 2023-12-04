package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getSolutionPart1(input string) int {
	rows := strings.Split(input, "\n")
	totalValue := 0
	for _, row := range rows {
		cardValue := 0
		r := strings.Split(row, ": ")

		var no int

		fmt.Sscanf(r[0], "Card %d", &no)

		cards := strings.Split(r[1], " | ")

		winningStrings := strings.Split(cards[0], " ")
		
		winning := make([]int, len(winningStrings))
		for i, s := range winningStrings {
			if len(s) > 0 {
				winning[i], _ = strconv.Atoi(s)
			}
		}

		myNumbersString := strings.Split(cards[1], " ")
		myNumbers := make(map[int]int)
		for _, s := range myNumbersString {
			if len(s) > 0 {
				v, _ := strconv.Atoi(s)
				myNumbers[v] = v
			}
		}

		for _, v := range winning {
			
			if _, ok := myNumbers[v]; ok {
				if cardValue == 0 {
					cardValue = 1
				} else {
					cardValue *= 2
				}
			}
		}
		totalValue += cardValue
	}
	return totalValue
}

func getSolutionPart2(input string) int {
	rows := strings.Split(input, "\n")
	max := len(rows)

	cardCopies := make(map[int]int)
	for _, row := range rows {
		cardValue := 0
		r := strings.Split(row, ": ")

		var no int

		fmt.Sscanf(r[0], "Card %d", &no)
		if val, ok := cardCopies[no]; ok {
			cardCopies[no] = val + 1
		} else {
			cardCopies[no] = 1
		}
		cards := strings.Split(r[1], " | ")

		winningStrings := strings.Split(cards[0], " ")
		
		winning := make([]int, len(winningStrings))
		for i, s := range winningStrings {
			if len(s) > 0 {
				winning[i], _ = strconv.Atoi(s)
			}
		}

		myNumbersString := strings.Split(cards[1], " ")
		myNumbers := make(map[int]int)
		for _, s := range myNumbersString {
			if len(s) > 0 {
				v, _ := strconv.Atoi(s)
				myNumbers[v] = v
			}
		}

		for _, v := range winning {
			
			if _, ok := myNumbers[v]; ok {
				cardValue++
			}
		}
		if cardValue > 0 {
			for x:=no+1; x<=no+cardValue; x++ {
				if(x > max) {
					break;
				}
				if val, ok := cardCopies[x]; ok {
					cardCopies[x] = val + (1 * cardCopies[no])
				} else {
					cardCopies[x] = (1 * cardCopies[no])
				}
			}
		}
		// totalValue += cardValue * cardCopies[no]
	}
	result := 0
	for _, v := range cardCopies {
		result += v
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
