package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

//go:embed input.txt
var input string

type Maps []Map

type Map struct {
	name string
	order int
	ranges []Range
}

type SeedRange struct {
	from int64
	to int64
}

type Range struct {
	destinationRangeStart int64
	sourceRangeStart int64
	rangeLength int64
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func (m Map) Convert(value int64) int64 {
	for _, r := range m.ranges {
		if value >= r.sourceRangeStart && value <= r.sourceRangeStart + r.rangeLength {
			offset := value - r.sourceRangeStart 
			return r.destinationRangeStart + offset
		}
	}
	return value
}

type Seed struct {
	value uint64
}

type AtoB struct {
	destinationRangeStart uint64
	sourceRangeStart      uint64
	rangeLength           uint64
}

var seeds []Seed

var mappers [][]AtoB = make([][]AtoB, 7)
var minLocation uint64 = math.MaxUint64

var mu sync.Mutex

var re = regexp.MustCompile("[0-9]+")

// Len is part of sort.Interface.
func (m Maps) Len() int {
	return len(m)
}

// Swap is part of sort.Interface.
func (m Maps) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (m Maps) Less(i, j int) bool {
	return m[i].order < m[j].order
}

func getSolutionPart1(input string) int64 {
	seeds := []int{}
	maps := Maps{}
	m := Map{}
	x := 0
	for i, row := range strings.Split(input, "\n") {
		if i == 0 {
			for _, seed := range strings.Split(strings.Split(row, ": ")[1], " ") {
				s, _ := strconv.Atoi(seed)
				seeds = append(seeds, s)
			}
			continue;
		}
		if len(row) == 0 {
			if len(m.ranges) > 0 {
				maps = append(maps, m)
			}
			
			m = Map{}
			m.order = x
			x++
			continue;
		} else {
			fields := strings.Split(row, " ")
			_, ok := strconv.Atoi(fields[0])
			if ok == nil {
				r := Range{}
				for i, fs := range fields {
					if v, o := strconv.Atoi(fs); o == nil {
						switch i {
						case 0:
							r.destinationRangeStart = int64(v)
						case 1:
							r.sourceRangeStart = int64(v)
						case 2:
							r.rangeLength = int64(v)
						}
					}
				}
				m.ranges = append(m.ranges, r)
			} else {
				m.name = fields[0]
			}
		}
	}
	if len(m.ranges) > 0 {
		maps = append(maps, m)
	}

	sort.Sort(maps)

	locations := []int64{}
	for _, s := range seeds {
		v := int64(s)
		for _, m := range maps {
			v = m.Convert(v)
		}

		locations = append(locations, v)
	}

	sort.Slice(locations, func(i, j int) bool { return locations[i] < locations[j] })

	return int64(locations[0])
}




func getSolutionPart2(input string) int64 {
	split := strings.Split(input, "\n\n")

	seedsStr := split[0]
	seedsNumbersStr := re.FindAllString(seedsStr, -1)

	for _, seedNumberStr := range seedsNumbersStr {
		seedNumber, err := strconv.Atoi(seedNumberStr)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, Seed{uint64(seedNumber)})
	}

	var seedsPairs [][]Seed
	for i := 0; i < len(seeds); i += 2 {
		seedsPairs = append(seedsPairs, []Seed{seeds[i], seeds[i+1]})
	}

	for index, split := range split[1:] {
		buildMap(split, index+1)
	}

	var wg sync.WaitGroup
	wg.Add(len(seedsPairs))

	for _, seedPair := range seedsPairs {
		fmt.Println("Calculating for seed pair", seedPair[0].value, seedPair[1].value)
		go runForSeedPair(seedPair, &wg)
	}

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Min location is", minLocation)
	fmt.Println("Done!")
	return int64(minLocation)
}

func runForSeedPair(seedPair []Seed, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := seedPair[0].value; i <= seedPair[0].value+seedPair[1].value-1; i++ {
		calculateForSource(Seed{i}.value, 0)
	}
}

func calculateForSource(source uint64, index uint64) {
	nextSource := source
	for _, mapper := range mappers[index] {
		if isBetween(source, mapper.sourceRangeStart, mapper.sourceRangeStart+mapper.rangeLength-1) {
			nextSource = source + (mapper.destinationRangeStart - mapper.sourceRangeStart)
		}
	}

	if index < 6 {
		calculateForSource(nextSource, index+1)
	} else {
		if nextSource < minLocation {
			mu.Lock()
			minLocation = nextSource
			mu.Unlock()
		}
	}
}

func buildMap(split string, index int) {
	elementsWithoutTitle := strings.Split(split, "\n")[1:]
	build(elementsWithoutTitle, &mappers[index-1])
}

func build(lines []string, mapToUpdate *[]AtoB) {
	for _, line := range lines {
		elementStr := re.FindAllString(line, -1)

		destinationRangeStart, err := strconv.Atoi(elementStr[0])
		if err != nil {
			panic(err)
		}
		sourceRangeStart, err := strconv.Atoi(elementStr[1])
		if err != nil {
			panic(err)
		}
		rangeLength, err := strconv.Atoi(elementStr[2])
		if err != nil {
			panic(err)
		}

		*mapToUpdate = append(*mapToUpdate, AtoB{uint64(destinationRangeStart), uint64(sourceRangeStart), uint64(rangeLength)})
	}
}

// 
func isBetween(num, min, max uint64) bool {
	return num >= min && num <= max


}

func main() {
	if os.Getenv("part") == "part2" {
		//fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}
}
