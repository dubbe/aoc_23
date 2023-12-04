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

func getAdjacentStars(grid map[Point]string, point NumberRange) []Point{

	points := []Point{}
	for x := point.fromX -1; x<=point.toX +1; x++ {
		for y := point.y -1; y<=point.y +1; y++ {
			if x >= 0 && y >= 0 {
				if _, ok := grid[Point{x, y}]; ok {
					// fmt.Printf("found: %s \n", val)
					points = append(points, Point{x, y})
				}
			}
		}
	}
	return points
}



func getSolutionPart1(input string) int64 {
	result := int64(0)
	grid := []NumberRange{}
	aSymbol := make(map[Point]string)
	for y, row := range strings.Split(input, "\n") {
		isNumber := false;
		number := ""
		numberFromX := 0

		
		for x, v := range row {
			value := string(v)
			if _, err := strconv.Atoi(value); err == nil {
				if !isNumber {
					numberFromX = x
				}
				isNumber = true
				number += value
				if len(row) == x + 1 {
					numberRange := NumberRange{numberFromX, x, y, number}
					grid = append(grid, numberRange)
					number = ""
				}
			} else  {
				if isNumber {
					numberRange := NumberRange{numberFromX, x-1, y, number}
					grid = append(grid, numberRange)
					number = ""
				}
				isNumber = false
				
				if value != "." {
					aSymbol[Point{x: x, y: y}] = value
				}
			} 
			
		}

	}

	for _, point := range grid {
		r := hasAdjacent(aSymbol, point)
		if r {
			if i, ok := strconv.Atoi(point.value); ok == nil {		
				result += int64(i)
			}
		}
	}

	return result
}

func getSolutionPart2(input string) int64 {
	result := int64(0)
	grid := []NumberRange{}
	aSymbol := make(map[Point]string)
	for y, row := range strings.Split(input, "\n") {
		isNumber := false;
		number := ""
		numberFromX := 0

		
		for x, v := range row {
			value := string(v)
			if _, err := strconv.Atoi(value); err == nil {
				if !isNumber {
					numberFromX = x
				}
				isNumber = true
				number += value
				if len(row) == x + 1 {
					numberRange := NumberRange{numberFromX, x, y, number}
					grid = append(grid, numberRange)
					number = ""
				}
			} else  {
				if isNumber {
					numberRange := NumberRange{numberFromX, x-1, y, number}
					grid = append(grid, numberRange)
					number = ""
				}
				isNumber = false
				
				if value == "*" {
					aSymbol[Point{x: x, y: y}] = value
				}
			} 
			
		}

	}


	numbersWithGears := make(map[Point][]int)

	for _, point := range grid {
		v, _ := strconv.Atoi(point.value)
		r := getAdjacentStars(aSymbol, point)
		if len(r) > 0 {

				for _, p := range r {
					if val, ok := numbersWithGears[p]; ok {
						
						numbersWithGears[p] = append(val, v)
					} else {
						numbersWithGears[p] = []int{v}
					}
				}
			}
	}

	for _, v := range numbersWithGears {
		if len(v) == 2 {
			result += int64(v[0] * v[1])
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
