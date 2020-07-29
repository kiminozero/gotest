package leetcode

import (
    "strings"
)
func lengthOfLongestSubstring(s string) int {
	var Length int
	var s1 string
	left := 0
	right := 0
	s1 = s[left:right]
	for ; right < len(s); right++ {

		if index := strings.IndexByte(s1, s[right]); index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > Length {
			Length = len(s1)
		}
	}

	return Length
}
