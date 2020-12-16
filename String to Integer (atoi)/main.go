package main

import "math"

func main() {

}

func myAtoi(s string) int {
	t := 1
	sum := 0
	plu := false
	for _, value := range s {
		value1 := int(value)
		if value1 == 32 && plu == false {
			continue
		}
		if (value1 == 43 || value1 == 45) && plu == false {
			plu = true
			if value1 == 45 {
				t = -1
			}
			continue
		}
		if value1 >= 48 && value1 <= 57 {
			sum = sum*10 + value1 - 48
			plu = true
			if sum*t > math.MaxInt32 {
				return math.MaxInt32
			} else if sum*t < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}
	return sum * t
}

//45 43  48~57
// Input: str = "-91283472332"
// Output: -2147483648
// Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer. Thefore INT_MIN (−231) is returned.
