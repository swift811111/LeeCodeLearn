package main

import "strconv"

func main() {

}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return csrs(countAndSay(n - 1))
}

func csrs(s string) string {
	num := 0
	str := ""
	for i := 0; i < len(s); i++ {
		num++
		if i == len(s)-1 || s[i] != s[i+1] {
			n, _ := strconv.Atoi(string(s[i]))
			str = str + strconv.Itoa(num) + strconv.Itoa(n)
			num = 0
		}
	}
	return str
}

// Input: n = 4
// Output: "1211"
// Explanation:
// countAndSay(1) = "1"
// countAndSay(2) = say "1" = one 1 = "11"
// countAndSay(3) = say "11" = two 1's = "21"
// countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"
