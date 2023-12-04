package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed input-test.txt
var testinput string


func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := int64(4361)
	actualSolution := getSolutionPart1(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestHasAdjacent(t *testing.T) {
	number := NumberRange{1,1,1,"1"}
	test := []string{"-", "&", "+", "*", "#", "/", "="}
	for _, te := range test {
		points := make(map[Point]string)
		for x:=0;x<=2;x++ {
			for y:=0;y<=2;y++ {
				points[Point{x,y}] = te
				assert.True(t, hasAdjacent(points, number))
			}
		}
		
		
	}
}

// func TestAOC_getSolutionPart1Correct(t *testing.T) {
// 	expectedSolution := int64(550934)
// 	actualSolution := getSolutionPart1(input)
// 	assert.Equal(t, expectedSolution, actualSolution)
// }


func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := 0
	actualSolution := getSolutionPart2(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getSolutionPart1(testinput)
	}
}
