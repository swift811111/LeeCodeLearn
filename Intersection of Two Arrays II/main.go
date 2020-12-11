package main

func main() {

}

func intersect(nums1 []int, nums2 []int) []int {
	map2 := make(map[int]int, len(nums2))
	result := make([]int, 0, 5)
	for _, value := range nums2 {
		map2[value]++
	}
	for _, n1value := range nums1 {
		if map2[n1value] > 0 {
			result = append(result, n1value)
			map2[n1value]--
		}
	}
	return result
}

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]
