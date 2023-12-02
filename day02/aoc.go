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

	for _, game := range rows {
		r := strings.Split(game, ":")
		id, _ := strconv.Atoi(strings.Split(r[0], " ")[1])

		possible := true;
		for _, t := range strings.Split(r[1][1:], "; ") {
			cubes := strings.Split(t, ", ")

			for _, cube := range cubes {
				c := strings.Split(cube, " ")
				color := c[1]
				number, _ := strconv.Atoi(c[0])

				if (color == "red" && number > 12) || (color == "green" && number > 13) || (color == "blue" && number > 14) {
					possible = false;
					break;
				}
			}
			
		}
		if(possible) {
			result += id
		}
	}
	return result
}

func getSolutionPart2(input string) int {
	result := 0
	rows := strings.Split(input, "\n")

	for _, game := range rows {
		r := strings.Split(game, ":")
		drawnCubes := make(map[string]int)

		for _, t := range strings.Split(r[1][1:], "; ") {
			cubes := strings.Split(t, ", ")

			for _, cube := range cubes {
				c := strings.Split(cube, " ")
				color := c[1]
				number, _ := strconv.Atoi(c[0])
				val, ok := drawnCubes[color]
				if ok {
					if number > val {
						drawnCubes[color] = number
					} 
				} else {
					drawnCubes[color] = number
				}
			}
			
		}

		gameSum := 1;
		for _, k := range drawnCubes {
			gameSum = gameSum * k
		}
		result += gameSum

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
