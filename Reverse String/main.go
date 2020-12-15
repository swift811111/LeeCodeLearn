package main

func main() {

}

func reverseString(s []byte) {
	v := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		v = append(v, s[i])
	}
	for key, value := range v {
		s[key] = value
	}
}

// Input: ["H","a","n","n","a","h"]
// Output: ["h","a","n","n","a","H"]
