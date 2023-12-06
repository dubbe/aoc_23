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

type Race struct {
	Time     uint64
	Distance uint64
}

func getSolutionPart1(input string) uint64 {
	rows := strings.Split(input, "\n")
	races := []Race{}
	re := regexp.MustCompile("[0-9]+")
	times := re.FindAllString(rows[0], -1)
	distances := re.FindAllString(rows[1], -1)

	fmt.Println(times, distances)
	for i, _ := range times {
		time, _ := strconv.ParseUint(times[i], 10, 64)
		distance, _ := strconv.ParseUint(distances[i], 10, 64)
		races = append(races, Race{time, distance})
	}

	noRecordsBeaten := uint64(0)
	for _, race := range races {
		raceBeaten := []uint64{}
		for x := uint64(0); x <= race.Time; x++ {
			speed := 1 * x
			distance := speed * (race.Time - x)
			if distance > race.Distance {
				raceBeaten = append(raceBeaten, x)
			}
		}
		if noRecordsBeaten == 0 {
			noRecordsBeaten = uint64(len(raceBeaten))
		} else {
			noRecordsBeaten *= uint64(len(raceBeaten))
		}
	}

	return noRecordsBeaten
}

func getSolutionPart2(input string) uint64 {

	rows := strings.Split(input, "\n")
	re := regexp.MustCompile("[0-9]+")
	times := re.FindAllString(rows[0], -1)
	distances := re.FindAllString(rows[1], -1)

	time := ""
	distance := ""

	for i, _ := range times {
		time += times[i]
		distance += distances[i]
	}

	t, _ := strconv.ParseUint(time, 10, 64)
	d, _ := strconv.ParseUint(distance, 10, 64)
	noRecordsBeaten := uint64(0)

	for x := uint64(0); x <= t; x++ {
		speed := 1 * x
		distance := speed * (t - x)
		if distance > d {
			noRecordsBeaten++
		}
	}

	return noRecordsBeaten
}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart2(input))
	}
}
