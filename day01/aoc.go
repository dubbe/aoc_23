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
	result := 0
	rows := strings.Split(input, "\n")
	for _, r := range rows {
		fmt.Println(r)
		newString := ""
		for _, d := range r {
			i, err := strconv.Atoi(string(d))
			if err == nil {
				newString += fmt.Sprint(i)
			}
		}	

		first3, _ := strconv.Atoi(fmt.Sprintf("%s%s", newString[0:1], newString[len(newString)-1:]))

		result = result + first3
	}
	return result
}

func getSolutionPart2(input string) int {
	numbers := map[string]string{"one":"1", "two":"2", "three":"3", "four":"4", "five":"5", "six":"6", "seven":"7", "eight":"8", "nine":"9"}
	rows := strings.Split(input, "\n")
	newRows := ""
	for _, r := range rows {
		newRow := r
		for k, v := range numbers {
			newRow = strings.Replace(newRow, k, fmt.Sprintf("%s%s%s",k, v, k), -1)
			fmt.Printf("%s, %s = %s\n", k, r, newRow)
		}
		// fmt.Printf("%s = %s\n",r, newRow)
		newRows = fmt.Sprintf("%s\n%s", newRows, newRow)
		
	}

	return getSolutionPart1(newRows[1:])
}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
