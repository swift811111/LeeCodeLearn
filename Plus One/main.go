package main

func main() {

}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 > 9 {
			digits[i] = 0
			if i == 0 {
				temp := append([]int{1}, digits...)
				return temp
			}
		} else {
			digits[i] += 1
			return digits
		}
	}
	return digits
}

// Input: digits = [4,3,2,1]
// Output: [4,3,2,2]
// Explanation: The array represents the integer 4321.
