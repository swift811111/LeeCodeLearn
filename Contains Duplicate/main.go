package main

func main() {

}

func containsDuplicate(nums []int) bool {
	dupMap := make(map[int]int, 4)
	for _, value := range nums {
		dupMap[value]++
		if dupMap[value] >= 2 {
			return true
		}
	}
	return false
}

// Input: [1,1,1,3,3,4,3,2,4,2]
// Output: true

// Input: [1,2,3,4]
// Output: false
