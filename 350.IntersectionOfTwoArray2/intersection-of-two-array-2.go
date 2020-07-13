package IntersectionOfTwoArray2

func intersect(nums1 []int, nums2 []int) []int {
	if nums2 == nil || nums1 == nil {
		return nil
	}

	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	if len(nums1) < len(nums2) {
		tem := nums1
		nums1 = nums2
		nums2 = tem
	}
	hashMap := map[int]int{}
	var intersection []int

	for _, v := range nums2 {
		hashMap[v]++
	}

	for _, v := range nums1 {
		if hashMap[v] > 0 {
			intersection = append(intersection, v)
			hashMap[v]--
		}
	}
	return intersection
}
