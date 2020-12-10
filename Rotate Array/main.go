package main

func main() {

}

func rotate(nums []int, k int) {
	remain := k % len(nums)
	temp := make([]int, 0, len(nums))
	temp = append(nums[len(nums)-remain:], nums[:len(nums)-remain]...)
	for key, value := range temp {
		nums[key] = value
	}
}

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
