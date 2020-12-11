package main

func main() {

}

func moveZeroes(nums []int) {
	index := 0
	for key, value := range nums {
		if value != 0 {
			nums[key] = 0
			nums[index] = value
			index++
		}
	}
}

// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]
