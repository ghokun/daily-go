package dailygo

/*
Given an array of integers, find the first missing positive integer in linear
time and constant space. In other words, find the lowest positive integer that
does not exist in the array. The array can contain duplicates and negative
numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place.
*/

func _20220315(arr []int) int {
	moveToValuePosition := func(arr []int, index int, value int) int {
		temp := arr[value-1]
		arr[value-1] = arr[index]
		arr[index] = temp
		return temp
	}

	length := len(arr)
	for index, value := range arr {
		// first pass
		if value <= 0 || value > length {
			continue
		}
		if arr[value-1] == value {
			continue
		}
		temp := moveToValuePosition(arr, index, value)
		// first pass
		if temp <= 0 || temp > length {
			continue
		}
		if arr[temp-1] == temp {
			continue
		}
		moveToValuePosition(arr, index, temp)
	}
	for index, value := range arr {
		if value != index+1 {
			return index + 1
		}
	}
	return length + 1
}
