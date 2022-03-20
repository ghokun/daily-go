package dailygo

import "strconv"

/*
Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.

For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.

You can assume that the messages are decodable. For example, '001' is not allowed.
*/
func _20220317(message string) int {
	if len(message) == 0 {
		return 0
	}
	if len(message) == 1 {
		return 1
	}
	count1 := _20220317(message[1:])
	count2 := 0
	first2chars := message[:2]
	firstTwo, err := strconv.Atoi(first2chars)
	if err == nil && firstTwo >= 10 && firstTwo <= 26 {
		if len(message) == 2 {
			return 2
		}
		count2 = _20220317(message[2:])
	}
	return count1 + count2
}
