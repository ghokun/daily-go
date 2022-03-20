package dailygo

/*
Given an array of integers, return a new array such that each element at index i
of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be
[120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
*/
func _20220313(numbers []int) []int {
	length := len(numbers)
	var forward = make([]int, length)
	var backward = make([]int, length)
	var product = make([]int, length)

	forward[0] = 1
	backward[length-1] = 1

	for i := 1; i < length; i++ {
		forward[i] = forward[i-1] * numbers[i-1]
		backward[length-i-1] = backward[length-i] * numbers[length-i]
	}
	for i := 0; i < length; i++ {
		product[i] = forward[i] * backward[i]
	}
	return product
}
