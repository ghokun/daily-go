package dailygo

/*
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/
func _20220312(number int, numbers []int) bool {
	var numberMap = make(map[int]int)
	for _, num := range numbers {
		diff := number - num
		if _, ok := numberMap[diff]; ok {
			return true
		}
		numberMap[num] = diff
	}
	return false
}
