package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getSolutionPart1(input string) int64 {
	rows := strings.Split(input, "\n")

	re := regexp.MustCompile("-?[0-9]+")

	result := int64(0)

	for _, row := range rows {
		if row == "" {
			continue
		}
		numberSeries := [][]int64{}
		fmt.Printf("\n\nRow: %s\n", row)
		numbersString := re.FindAllString(row, -1)
		numbers := []int64{}
		for _, ns := range numbersString {
			n, _ := strconv.ParseInt(ns, 10, 64)
			numbers = append(numbers, n)
		}
		numberSeries = append(numberSeries, numbers)
		found := false
		for !found {
			zeroes := 0
			newNumbers := []int64{}
			for i := range numbers {
				if i == len(numbers)-1 {
					break
				}
				difference := numbers[i+1] - numbers[i]
				newNumbers = append(newNumbers, difference)
				if difference == 0 {
					zeroes++
				}

			}

			if zeroes == len(newNumbers) {
				found = true
			}
			numbers = newNumbers
			fmt.Println(numbers)
			numberSeries = append(numberSeries, numbers)
		}

		prev := int64(0)
		for i := len(numberSeries) - 2; i >= 0; i-- {
			last := numberSeries[i][len(numberSeries[i])-1]
			prev = last + prev
			if i == 0 {
				fmt.Printf("Solution: %d\n", prev)
				result += prev
			}
		}
	}

	return result
}

func getSolutionPart2(input string) int64 {
	rows := strings.Split(input, "\n")

	re := regexp.MustCompile("-?[0-9]+")

	result := int64(0)

	for _, row := range rows {
		if row == "" {
			continue
		}
		numberSeries := [][]int64{}
		fmt.Printf("\n\nRow: %s\n", row)
		numbersString := re.FindAllString(row, -1)
		numbers := []int64{}
		for _, ns := range numbersString {
			n, _ := strconv.ParseInt(ns, 10, 64)
			numbers = append(numbers, n)
		}
		numberSeries = append(numberSeries, numbers)
		found := false
		for !found {
			zeroes := 0
			newNumbers := []int64{}
			for i := range numbers {
				if i == len(numbers)-1 {
					break
				}
				difference := numbers[i+1] - numbers[i]
				newNumbers = append(newNumbers, difference)
				if difference == 0 {
					zeroes++
				}

			}

			if zeroes == len(newNumbers) {
				found = true
			}
			numbers = newNumbers
			fmt.Println(numbers)
			numberSeries = append(numberSeries, numbers)
		}

		prev := int64(0)
		for i := len(numberSeries) - 2; i >= 0; i-- {
			first := numberSeries[i][0]
			prev = first - prev
			if i == 0 {
				fmt.Printf("Solution: %d\n", prev)
				result += prev
			}
		}
	}

	return result
}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart2(input))
	}
}
