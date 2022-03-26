package dailygo

import (
	"testing"
)

func Test_20220326(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	distribution := make(map[int]float64)

	categorySize := len(elements)
	totalFreq := 1000
	for i := 0; i < totalFreq; i++ {
		distribution[_20220326(elements)]++
	}
	expectedFreq := float64(totalFreq / categorySize)
	chi := 0.0
	for _, observedFreq := range distribution {
		chi = chi + (((observedFreq - expectedFreq) * (observedFreq - expectedFreq)) / expectedFreq)
	}
	// df = 9
	// 0.05 = 16.919
	if chi > 16.919 {
		t.Errorf("chi is %f, which is greater than 16.919", chi)
	}
}
