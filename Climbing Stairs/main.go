package main

func main() {

}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	n1, n2 := 1, 2

	for i := 3; i <= n; i++ {
		temp := n1 + n2
		n1 = n2
		n2 = temp
	}
	return n2
}

// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
