package main

import "math"

func main() {

}

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	t := 1
	if x < 0 {
		x *= -1
		t = -1
	}
	sum := 0
	for x/10 != 0 || x%10 != 0 {
		sum = sum*10 + x%10
		x /= 10
	}
	if sum*t > math.MaxInt32 || sum*t < math.MinInt32 {
		return 0
	}
	return sum * t
}

// Input: x = -123
// Output: -321

// Input: x = 120
// Output: 21
