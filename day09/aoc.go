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

var re = regexp.MustCompile("-?[0-9]+")

func getNumbersSeries(input string) [][]int64 {

	numberSeries := [][]int64{}
	numbersString := re.FindAllString(input, -1)
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
	return numberSeries
}

func getSolutionPart1(input string) int64 {
	rows := strings.Split(input, "\n")

	result := int64(0)

	for _, row := range rows {
		if row == "" {
			continue
		}
		numberSeries := getNumbersSeries(row)

		prev := int64(0)
		for i := len(numberSeries) - 2; i >= 0; i-- {
			last := numberSeries[i][len(numberSeries[i])-1]
			prev = last + prev
			if i == 0 {
				result += prev
			}
		}
	}

	return result
}

func getSolutionPart2(input string) int64 {
	result := int64(0)

	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}
		numberSeries := getNumbersSeries(row)

		prev := int64(0)
		for i := len(numberSeries) - 2; i >= 0; i-- {
			first := numberSeries[i][0]
			prev = first - prev
			if i == 0 {
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
		fmt.Println(getSolutionPart1(input))
	}
}
