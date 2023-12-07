package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed input-test.txt
var testinput string

func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := uint64(6440)
	actualSolution := getSolutionPart1(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := uint64(5905)
	actualSolution := getSolutionPart2(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getSolutionPart1(input)
	}
}
