package dailygo

/*
Given an integer k and a string s, find the length of the longest substring that
contains at most k distinct characters.

For example, given s = "abcba" and k = 2, the longest substring with k distinct characters is "bcb".
*/
func _20220324(str string, k int) string {
	longest := ""
	for i := 0; i < len(str); i++ {
		current := ""
		charMap := make(map[string]bool)
		for j := i; j < len(str); j++ {
			charMap[str[j:j+1]] = true
			if len(charMap) <= k {
				current += str[j : j+1]
			} else {
				break
			}
		}
		if len(current) > len(longest) {
			longest = current
		}
	}
	return longest
}
