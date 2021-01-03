package main

func main() {

}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	sum := 0
	low := prices[0]

	for _, value := range prices[1:] {
		if value > low {
			if value-low > sum {
				sum = value - low
			}
		}

		if value < low {
			low = value
		}
	}

	return sum
}

// Input: [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Not 7-1 = 6, as selling price needs to be larger than buying price.
