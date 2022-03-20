package dailygo

/*
Given a list of integers, write a function that returns the largest sum of
non-adjacent numbers. Numbers can be 0 or negative.

For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5]
should return 10, since we pick 5 and 5.

Follow-up: Can you do this in O(N) time and constant space?
*/
func _20220320(arr []int) int {
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 2 {
		return max(arr[0], arr[1])
	}
	zero := arr[0] + _20220320(arr[2:])
	one := arr[1] + _20220320(arr[3:])
	return max(zero, one)
}
