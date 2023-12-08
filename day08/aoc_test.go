package main

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed input-test.txt
var testinput string

//go:embed input-test-two.txt
var testinputTwo string

func TestAOC_getSolutionPart1(t *testing.T) {
	expectedSolution := uint64(6)
	actualSolution := getSolutionPart1(testinput)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestAOC_getSolutionPart2(t *testing.T) {
	expectedSolution := uint64(6)
	actualSolution := getSolutionPart2(testinputTwo)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestLCM(t *testing.T) {
	expectedSolution := uint64(6)
	actualSolution := lcm(17873, 17287)
	assert.Equal(t, expectedSolution, actualSolution)
}

func TestLCMMultple(t *testing.T) {
	expectedSolution := uint64(18625484023687)
	actualSolution := lcmMultiple([]uint64{17873, 17287, 23147, 19631, 13771, 20803})
	assert.Equal(t, expectedSolution, actualSolution)
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getSolutionPart1(input)
	}
}
