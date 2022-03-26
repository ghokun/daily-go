package dailygo

import (
	"math/rand"
	"time"
)

/*
Given a stream of elements too large to store in memory, pick a random element
from the stream with uniform probability.
*/
func _20220326(elements []int) int {
	rand.Seed(time.Now().UnixNano())
	var selected int
	for ndx, val := range elements {
		if rand.Intn(ndx+1)+1 == 1 {
			selected = val
		}
	}
	return selected
}
