package main

func main() {

}

func singleNumber(nums []int) int {
	result := 0
	for _, value := range nums {
		result ^= value
	}
	return result
}

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// Input: nums = [4,1,2,1,2]
// Output: 4

// 000
// 100
// 001
// 010
// 001
// 010
