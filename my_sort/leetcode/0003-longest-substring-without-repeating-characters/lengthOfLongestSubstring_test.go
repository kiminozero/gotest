package leetcode

import (
    "testing"
)

type testData struct {
    s       string
    length  int
}

func TestlengthOfLongestSubstring(t *testing.T) {
    tests := []testData{
        testData{"abcabcbb", 3},
        testData{"bbbbb", 1},
        testData{"pwwkew", 3},
    }

    for i, testcase := range tests {
        s := testcase.s
        length := testcase.length
        res := lengthOfLongestSubstring(s)
        if length != res {
            t.Errorf("lengthOfLongestSubstring failed at case %d, \n%s", i+1, s)
        }
    }
}
