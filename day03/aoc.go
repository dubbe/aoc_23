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

type Point struct {
	x int;
	y int;
}

type NumberRange struct {
	fromX int;
	toX int;
	y int;
	value string;
}

func hasAdjacent(grid map[Point]string, point NumberRange) bool{

	for x := point.fromX -1; x<=point.toX +1; x++ {
		for y := point.y -1; y<=point.y +1; y++ {
			if x >= 0 && y >= 0 {
				if _, ok := grid[Point{x, y}]; ok {
					// fmt.Printf("found: %s \n", val)
					return true
				}
			}
		}
	}
	return false
}

func getSolutionPart1(input string) int64 {
	result := int64(0)
	grid := []NumberRange{}
	aSymbol := make(map[Point]string)
	test := make(map[string]bool)
	for y, row := range strings.Split(input, "\n") {
		isNumber := false;
		number := ""
		numberFromX := 0

		
		for x, v := range row {
			value := string(v)
			test[value] = true
			if _, err := strconv.Atoi(value); err == nil {
				if !isNumber {
					numberFromX = x
				}
				isNumber = true
				number += value
			} else  {
				if isNumber {
					numberRange := NumberRange{numberFromX, x-1, y, number}
					grid = append(grid, numberRange)
					number = ""
				}
				isNumber = false
				
				if value != "." {
					test[value] = true
					aSymbol[Point{x: x, y: y}] = value
				}
			} 
			
		}

	}

	fmt.Println(test)

	for _, point := range grid {
		// fmt.Printf("checking %s\n", point.value)
		r := hasAdjacent(aSymbol, point)
		if r {
			if i, ok := strconv.Atoi(point.value); ok == nil {
				
				result += int64(i)
				// fmt.Printf("i: %d, result: %d\n", i, result)
			}
		}
	}

	return result
}

func getSolutionPart2(input string) int {
	result := 0
	
	return result
}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
