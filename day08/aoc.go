package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/TheAlgorithms/Go/math/gcd"
)

//go:embed input.txt
var input string

type Node struct {
	left  string
	right string
}

func getSolutionPart1(input string) uint64 {
	rows := strings.Split(input, "\n")
	instructions := rows[0]

	nodes := make(map[string]Node)

	for _, row := range rows[2:] {
		r := strings.Split(row, " = ")
		key := r[0]

		left := r[1][1:4]
		right := r[1][6:9]
		nodes[key] = Node{
			left:  left,
			right: right,
		}
	}

	result := uint64(0)
	newNode := "AAA"
	current := nodes[newNode]

	for continueLoop := true; continueLoop; {
		for _, i := range instructions {
			result++
			switch i {
			case 'L':
				newNode = current.left

			case 'R':
				newNode = current.right
			}
			current = nodes[newNode]
			if newNode == "ZZZ" {
				continueLoop = false
				break
			}
		}
	}

	return result
}

func getSolutionPart2(input string) uint64 {
	rows := strings.Split(input, "\n")
	instructions := rows[0]

	nodes := make(map[string]Node)

	for _, row := range rows[2:] {
		r := strings.Split(row, " = ")
		key := r[0]

		left := r[1][1:4]
		right := r[1][6:9]
		nodes[key] = Node{
			left:  left,
			right: right,
		}
	}

	startNodes := []string{}

	for k, _ := range nodes {
		if k[2] == 'A' {
			startNodes = append(startNodes, k)
		}
	}

	stepsToFirst := []uint64{}

	for _, startNode := range startNodes {

		newNode := ""
		steps := uint64(0)
		current := nodes[startNode]
		for continueLoop := true; continueLoop; {
			for _, i := range instructions {

				steps++
				switch i {
				case 'L':
					newNode = current.left

				case 'R':
					newNode = current.right
				}
				current = nodes[newNode]
				if newNode[2] == 'Z' {
					continueLoop = false
					stepsToFirst = append(stepsToFirst, steps)
					break
				}
			}
		}
	}

	return lcmMultiple(stepsToFirst)
}

// lcm of multiple numbers
func lcmMultiple(numbers []uint64) uint64 {
	lcmnum := int64(1)
	for {
		for _, number := range numbers {
			lcmnum = lcm(int64(lcmnum), int64(number))
		}
		fmt.Printf("LCM of %v is %d", numbers, lcmnum)
		return uint64(lcmnum)
	}
}

func lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(gcd.Iterative(a, b)))
}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
