package main

func main() {

}

func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	h := make(map[rune]int)
	for _, value := range s {
		h[value]++
	}
	for key, value := range s {
		if h[value] == 1 {
			return key
		}
	}
	return -1
}

// Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.
// s = "leetcode"
// return 0.

// s = "loveleetcode"
// return 2.
