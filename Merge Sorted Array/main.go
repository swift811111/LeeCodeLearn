func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	sr := make([]int, 0, m)
	im := 0
	in := 0
	for len(sr) < len(nums1) {
		if in >= n {
			sr = append(sr, nums1[im])
			im++
			continue
		}
		if im >= m {
			sr = append(sr, nums2[in])
			in++
			continue
		}

		if nums1[im] >= nums2[in] {
			sr = append(sr, nums2[in])
			in++
		} else {
			sr = append(sr, nums1[im])
			im++
		}
	}
	for key, value := range sr {
		nums1[key] = value
	}
}

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]