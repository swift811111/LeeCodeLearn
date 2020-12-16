package main

func main() {

}

func isPalindrome(s string) bool {
	sindex := 0
	index := len(s) - 1
	for i := 0; i < len(s); i++ {
		n1 := int(s[sindex])
		n2 := int(s[index])
		if n1 >= 97 {
			n1 -= 32
		}
		if n2 >= 97 {
			n2 -= 32
		}
		if !((n1 >= 48 && n1 <= 57) || (n1 >= 65 && n1 <= 90)) {
			sindex++
			continue
		}
		if !((n2 >= 48 && n2 <= 57) || (n2 >= 65 && n2 <= 90)) {
			index--
			continue
		}
		if n1 != n2 {
			return false
		}
		i++
		sindex++
		index--
	}
	return true
}

// Input: "A man, a plan, a canal: Panama"
// Output: true
