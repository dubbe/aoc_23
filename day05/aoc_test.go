package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed input-test.txt
var testinput string


func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := int64(35)
	actualSolution := getSolutionPart1(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestConvert(t *testing.T) {
	seedToSoilMap := Map{
		"seed",
		0,
		[]Range{{50, 98, 2}, {52, 50, 48}},
	}

	assert.Equal(t, int64(0), seedToSoilMap.Convert(0))
	assert.Equal(t, int64(52), seedToSoilMap.Convert(50))
	assert.Equal(t, int64(98), seedToSoilMap.Convert(96))
	assert.Equal(t, int64(50), seedToSoilMap.Convert(98))
}


// func TestOverlap(t *testing.T) {
// 	assert.True(t, Overlaps(SeedRange{0,10}, Range{0, 5, 11}))
// 	assert.False(t, Overlaps(SeedRange{0,10}, Range{0, 11, 11}))
// 	assert.True(t, Overlaps(SeedRange{15,20}, Range{0, 11, 11}))
// 	assert.False(t, Overlaps(SeedRange{100,150}, Range{0, 11, 11}))
// 	assert.False(t, Overlaps(SeedRange{98,99}, Range{52,50,48}))
// }

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := int64(46)
	actualSolution := getSolutionPart2(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getSolutionPart1(input)
	}
}
