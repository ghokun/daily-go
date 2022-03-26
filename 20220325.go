package dailygo

import (
	"math/rand"
)

func _20220325() float64 {

	numOfPointsInsideCircle := 0.0
	numOfPointsInsideSquare := 0.0
	var pi float64

	iter := 0
	max := 2_000_000
	for {
		x := rand.Float64()
		y := rand.Float64()

		d := x*x + y*y

		if d <= 1.0 {
			numOfPointsInsideCircle++
		}
		numOfPointsInsideSquare++

		area := numOfPointsInsideCircle / numOfPointsInsideSquare
		pi = area * 4
		iter++
		if iter == max {
			break
		}
	}
	piInt := int(pi * 1000.0)
	return float64(piInt) / 1000.0
}
