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


func getSolutionPart1(input string) int {
	result := 0
	rows := strings.Split(input, "\n")
	re := regexp.MustCompile("[0-9]")
	for _, r := range rows {
		if len(r) == 0 {
			break
		}
		numbers := re.FindAllString(r, -1)

		sum, _ := strconv.Atoi(fmt.Sprintf("%s%s", numbers[0], numbers[len(numbers)-1]))
		fmt.Println(sum)
		result = result + sum
	}
	return result
}

func getSolutionPart2(input string) int {
	numbers := map[string]string{"one":"1", "two":"2", "three":"3", "four":"4", "five":"5", "six":"6", "seven":"7", "eight":"8", "nine":"9"}
	rows := strings.Split(input, "\n")

	newLine := ""
	for _, row := range rows {
		for i, r  := range row {
			if _, err := strconv.Atoi(string(r)); err == nil {
				newLine += string(r)
			} else {
				for k, v := range numbers {
					if found := strings.HasPrefix(row[i:], k); found {
						newLine += v
					}
				}
			}
			
		}
		newLine += "\n"
	}

	return getSolutionPart1(newLine)
}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
